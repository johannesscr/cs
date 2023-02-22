package divideandconquer

import (
	"fmt"
	"testing"
)

func TestMakeMatrix(t *testing.T) {
	tt := []struct {
		name string
		n int
		out [][]int
	}{
		{
			n: 0,
			out: [][]int{},
		},
		{
			n: 1,
			out: [][]int{
				{0},
			},
		},
		{
			n: 2,
			out: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			n: 3,
			out: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := MakeMatrix(tc.n)
			s1 := fmt.Sprintf("%v", tc.out)
			s2 := fmt.Sprintf("%v", o)
			if s1 != s2 {
				t.Errorf("expected matrix %s got %s", s1, s2)
			}
		})
	}
}

func TestSquareMatrixMultiplication(t *testing.T) {
	tt := []struct {
		name string
		A [][]int
		B [][]int
		out [][]int
	}{
		{
			A: [][]int{
				{1, 2},
				{3, 4},
			},
			B: [][]int{
				{5, 6},
				{7, 8},
			},
			out: [][]int{
				{19, 22},
				{43, 50},
			},
		},
		{
			A: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			B: [][]int{
				{1, 2, 3},
				{4, 5, 4},
				{1, 2, 3},
			},
			out: [][]int{
				{12, 18, 20},
				{30, 45, 50},
				{48, 72, 80},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := SquareMatrixMultiplication(tc.A, tc.B)
			s1 := fmt.Sprintf("%v", tc.out)
			s2 := fmt.Sprintf("%v", o)
			if s1 != s2 {
				t.Errorf("expected matrix %s got %s", s1, s2)
			}
		})
	}
}

func BenchmarkSquareMatrixMultiplication(b *testing.B) {
	A := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 200},
	}
	B := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 200},
	}
	for i := 0; i < b.N; i++ {
		SquareMatrixMultiplication(A, B)
	}
}

func TestPartitionMatrix(t *testing.T) {
	tt := []struct {
		name string
		A [][]int
		A11 [][]int
		A12 [][]int
		A21 [][]int
		A22 [][]int
	}{
		{
			A: [][]int{
				{1, 2},
				{3, 4},
			},
			A11: [][]int{{1}},
			A12: [][]int{{2}},
			A21: [][]int{{3}},
			A22: [][]int{{4}},
		},
		{
			A: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			A11: [][]int{
				{1, 2},
				{5, 6},
			},
			A12: [][]int{
				{3, 4},
				{7, 8},
			},
			A21: [][]int{
				{9, 10},
				{13, 14},
			},
			A22: [][]int{
				{11, 12},
				{15, 16},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			a11, a12, a21, a22 := PartitionMatrix(tc.A)
			sa11 := fmt.Sprintf("%v", a11)
			sa12 := fmt.Sprintf("%v", a12)
			sa21 := fmt.Sprintf("%v", a21)
			sa22 := fmt.Sprintf("%v", a22)

			sA11 := fmt.Sprintf("%v", tc.A11)
			sA12 := fmt.Sprintf("%v", tc.A12)
			sA21 := fmt.Sprintf("%v", tc.A21)
			sA22 := fmt.Sprintf("%v", tc.A22)

			if sa11 != sA11 {
				t.Errorf("expected A11 %s got %s", sA11, sa11)
			}
			if sa12 != sA12 {
				t.Errorf("expected A12 %s got %s", sA12, sa12)
			}
			if sa21 != sA21 {
				t.Errorf("expected A21 %s got %s", sA21, sa21)
			}
			if sa22 != sA22 {
				t.Errorf("expected A22 %s got %s", sA22, sa22)
			}
		})
	}
}

func TestAddMatrix(t *testing.T) {

}

func TestCombineMatrixPartition(t *testing.T) {

}