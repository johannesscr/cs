package divideandconquer

import (
	"testing"
)

func TestFindMaxCrossingSubarray(t *testing.T) {
	tt := []struct{
		name string
		A []int
		low int
		mid int
		high int
		outLow int
		outHigh int
		outSum int
	}{
		{
			A: []int{9, 8, 1, 3, 2, 5, 9, 5, 10, 4, 5},
			low: 0,
			mid: 5,
			high: 11,
			outLow: 0,
			outHigh: 10,
			outSum: 61,
		},
		{
			A: []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
			low: 0,
			mid: 7,
			high: 15,
			outLow: 7, // starts at index
			outHigh: 10, // ends at index
			outSum: 43, // sum 18 + 20 + (-7) + 12
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			l, h, s := FindMaxCrossingSubarray(tc.A, tc.low, tc.mid, tc.high)
			if l != tc.outLow {
				t.Errorf("expected output low %d got %d", tc.outLow, l)
			}
			if h != tc.outHigh {
				t.Errorf("expected output high %d got %d", tc.outHigh, h)
			}
			if s != tc.outSum {
				t.Errorf("expected output sum %d got %d", tc.outSum, s)
			}
		})
	}
}

func TestFindMaxSubarray(t *testing.T) {
	tt := []struct {
		name string
		A []int
		low int
		high int
		outLow int
		outHigh int
		outSum int
	}{
		{
			A: []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
			low: 0,
			high: 15,
			outLow: 7, // starts at index
			outHigh: 10, // ends at index
			outSum: 43, // sum 18 + 20 + (-7) + 12
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			l, h, s := FindMaxSubarray(tc.A, tc.low, tc.high)
			if l != tc.outLow {
				t.Errorf("expected output low %d got %d", tc.outLow, l)
			}
			if h != tc.outHigh {
				t.Errorf("expected output high %d got %d", tc.outHigh, h)
			}
			if s != tc.outSum {
				t.Errorf("expected output sum %d got %d", tc.outSum, s)
			}
		})
	}
}

func BenchmarkFindMaxSubarray(b *testing.B) {
	A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	for i := 0; i < b.N; i++ {
		FindMaxSubarray(A, 0, 15)
	}
}

func TestMaxSubarray(t *testing.T) {
	tt := []struct {
		name string
		A []int
		outLow int
		outHigh int
		outSum int
	}{
		{
			A: []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
			outLow: 7, // starts at index
			outHigh: 10, // ends at index
			outSum: 43, // sum 18 + 20 + (-7) + 12
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			l, h, s := MaxSubarray(tc.A)
			if l != tc.outLow {
				t.Errorf("expected output low %d got %d", tc.outLow, l)
			}
			if h != tc.outHigh {
				t.Errorf("expected output high %d got %d", tc.outHigh, h)
			}
			if s != tc.outSum {
				t.Errorf("expected output sum %d got %d", tc.outSum, s)
			}
		})
	}
}

func BenchmarkMaxSubarray(b *testing.B) {
	A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	for i := 0; i < b.N; i++ {
		MaxSubarray(A)
	}
}