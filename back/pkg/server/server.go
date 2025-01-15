package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/tprifti/gs/pkg/pack"
)

type Server struct {
	listenAddr string
	packSizes  []int
}

func NewServer(listenAddr string, packSizes []int) *Server {
	return &Server{
		listenAddr: listenAddr,
		packSizes:  packSizes,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/calculate", s.handleCalculatePackage)
	mux.HandleFunc("/packs", s.handleGetPackSizes)
	slog.Info("Starting server on", "address", s.listenAddr)
	go func() {
		err := http.ListenAndServe(s.listenAddr, withCORS(mux))
		if err != nil {
			slog.Error("Error starting server", "error", err)
		}
	}()
}

func (s *Server) handleGetPackSizes(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.handleUpdatePackSizes(w, r)
		return
	}
	WriteJSON(w, http.StatusOK, s.packSizes)
}

type PackSizeParams struct {
	PackSizes []int `json:"packSizes"`
}

func (s *Server) handleUpdatePackSizes(w http.ResponseWriter, r *http.Request) {
	var params PackSizeParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	s.packSizes = params.PackSizes
	WriteJSON(w, http.StatusOK, map[string]string{"message": "pack sizes updated"})
}

func (s *Server) handleCalculatePackage(w http.ResponseWriter, r *http.Request) {
	items := r.URL.Query().Get("items")
	itemsInt, err := strconv.Atoi(items)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "item query param is required"})
		return
	}

	if itemsInt <= 0 {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "item query param must be greater than 0"})
		return
	}

	packages := pack.CalculatePackages(s.packSizes, itemsInt)
	WriteJSON(w, http.StatusOK, packages)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-auth-token")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}
