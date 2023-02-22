package algorithms

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	in := []int{3, 2, 5, 6, 1, 3, 2, 6, 8, 9, 7}
	out := []int{1, 2, 2, 3, 3, 5, 6, 6, 7, 8, 9}
	o := BubbleSort(in)
	s1 := fmt.Sprintf("%v", out)
	s2 := fmt.Sprintf("%v", o)
	if s1 != s2 {
		t.Errorf("expected %s got %s", s1, s2)
	}
}

func TestSelectionSort(t *testing.T) {
	in := []int{3, 2, 5, 6, 1, 3, 2, 6, 8, 9, 7}
	out := []int{1, 2, 2, 3, 3, 5, 6, 6, 7, 8, 9}
	o := SelectionSort(in)
	s1 := fmt.Sprintf("%v", out)
	s2 := fmt.Sprintf("%v", o)
	if s1 != s2 {
		t.Errorf("expected %s got %s", s1, s2)
	}
}

func TestInsertionSort(t *testing.T) {
	in := []int{3, 2, 5, 6, 1, 3, 2, 6, 8, 9, 7}
	out := []int{1, 2, 2, 3, 3, 5, 6, 6, 7, 8, 9}
	o := InsertionSort(in)
	s1 := fmt.Sprintf("%v", out)
	s2 := fmt.Sprintf("%v", o)
	if s1 != s2 {
		t.Errorf("expected %s got %s", s1, s2)
	}
}

func TestMergeSort(t *testing.T) {
	in := []int{3, 2, 5, 6, 1, 3, 2, 6, 8, 9, 7}
	out := []int{1, 2, 2, 3, 3, 5, 6, 6, 7, 8, 9}
	o := MergeSort(in)
	s1 := fmt.Sprintf("%v", out)
	s2 := fmt.Sprintf("%v", o)
	if s1 != s2 {
		t.Errorf("expected %s got %s", s1, s2)
	}
}
