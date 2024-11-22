package main

import "testing"

func TestMaxTheStock(t *testing.T) {
	tests := map[string]struct {
		prices []int
		want   int
	}{
		"profit available": {
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		"no profit available": {
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		"big profit!": {
			prices: []int{1000, 45, 2000, 1975, 35, 500},
			want:   1955,
		},
		"max profit early with later minimum": {
			prices: []int{10, 500, 5, 100, 2, 300},
			want:   490,
		},
		"no data": {
			want: 0,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := maxTheStock(testCase.prices)
			if want := testCase.want; got != want {
				t.Errorf("got %d, but expected %d", got, want)
			}
		})
	}
}
