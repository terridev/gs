package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/tprifti/gs/pkg/server"
)

func main() {
	defaultPackSizes := []int{250, 500, 1000, 2000, 5000}
	server := server.NewServer(":8080", defaultPackSizes)
	server.Start()

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	<-sigch
	slog.Info("Server shutdown")
	os.Exit(0)
}
