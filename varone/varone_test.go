package varone

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name     string
		numbs    []int
		expected int
	}{
		{"pos", []int{3, 4, 5}, 12},
		{"neg", []int{5, -9, 1}, -3},
		{"empty", []int{}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := sum(tc.numbs...)
			if res != tc.expected {
				fmt.Errorf("expected %d, got %d", tc.expected, res)
			}
		})
	}
}
