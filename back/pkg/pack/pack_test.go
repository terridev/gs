package pack

import (
	"testing"
)

func TestCalculatePackages(t *testing.T) {
	defaultPackSizes := []int{250, 500, 1000, 2000, 5000}

	tests := []struct {
		items    int
		expected map[int]int
	}{
		{1, map[int]int{250: 1}},
		{250, map[int]int{250: 1}},
		{251, map[int]int{500: 1}},
		{750, map[int]int{500: 1, 250: 1}},
		{12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
		{499, map[int]int{500: 1}},
		{999, map[int]int{1000: 1}},
	}

	for _, test := range tests {
		result := CalculatePackages(defaultPackSizes, test.items)
		for _, pack := range result {
			if pack.Count != test.expected[pack.PackSize] {
				t.Errorf("pack size config: %v, total items: %d, expected %d x %d, got %d x %d", defaultPackSizes, test.items, test.expected[pack.PackSize], pack.PackSize, pack.Count, pack.PackSize)
			}
		}
	}
}
