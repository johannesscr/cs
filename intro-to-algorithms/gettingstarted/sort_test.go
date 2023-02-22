package gettingstarted

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8}
	output := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	o := InsertionSort(input)

	sfmt1 := fmt.Sprintf("%v", output)
	sfmt2 := fmt.Sprintf("%v", o)
	if sfmt1 != sfmt2 {
		t.Errorf("expected %v got %v", sfmt1, sfmt2)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8}
	for i := 0; i < b.N; i++ {
		InsertionSort(input)
	}
}
func BenchmarkInsertionSort2(b *testing.B) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8}
	for i := 0; i < b.N; i++ {
		InsertionSort(input)
	}
}

func TestSelectionSort(t *testing.T) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8}
	output := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	o := SelectionSort(input)

	s1 := fmt.Sprintf("%v", output)
	s2 := fmt.Sprintf("%v", o)
	if s1 != s2 {
		t.Errorf("expected %v got %v", s1, s2)
	}
}


func BenchmarkSelectionSort(b *testing.B) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8}
	for i := 0; i < b.N; i++ {
		SelectionSort(input)
	}
}
func BenchmarkSelectionSort2(b *testing.B) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8}
	for i := 0; i < b.N; i++ {
		SelectionSort(input)
	}
}

func TestMerge(t *testing.T) {
	tt := []struct{
		A []int
		p int
		q int
		r int
		output []int
	}{
		{
			A: []int{3, 2},
			p: 0,
			q: 0,
			r: 1,
			output: []int{2, 3},
		},
		{
			A: []int{1, 3, 2},
			p: 1,
			q: 1,
			r: 2,
			output: []int{1, 2, 3},
		},
		{
			A: []int{1, 3, 2, 4, 5, 7, 3, 6, 8, 9, 10},
			p: 3,
			q: 5,
			r: 10,
			output: []int{1, 3, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			o := Merge(tc.A, tc.p, tc.q, tc.r)

			s1 := fmt.Sprintf("%v", tc.output)
			s2 := fmt.Sprintf("%v", o)
			if s1 != s2 {
				t.Errorf("expected %v got %v", s1, s2)
			}
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	A := []int{1, 3, 2, 4, 5, 7, 3, 6, 8, 9, 10}
	p := 3
	q := 5
	r := 10

	for i := 0; i < b.N; i++ {
		Merge(A, p, q, r)
	}
}

func TestMergeSort(t *testing.T) {
	tt := []struct {
		name string
		A []int
		output []int
	}{
		{
			name: "",
			A: []int{1, 3, 2, 4, 5, 7, 3, 6, 8, 9, 10},
			output: []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "",
			A: []int{1, 3, 2, 5, 6, 7, 4, 8, 3, 5, 9, 2, 6, 7, 1, 2, 1, 4, 5, 10},
			output: []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 7, 7, 8, 9, 10},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := MergeSort(tc.A, 0, len(tc.A)-1)

			s1 := fmt.Sprintf("%v", tc.output)
			s2 := fmt.Sprintf("%v", o)
			if s1 != s2 {
				t.Errorf("expected output '%s' got '%s", s1, s2)
			}
		})
	}

}

func BenchmarkMergeSort(b *testing.B) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8}
	for i := 0; i < b.N; i++ {
		MergeSort(input, 0, len(input)-1)
	}
}

func BenchmarkMergeSor2(b *testing.B) {
	input := []int{9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8, 9, 3, 1, 4, 5, 6, 7, 2, 8}
	for i := 0; i < b.N; i++ {
		MergeSort(input, 0, len(input)-1)
	}
}
