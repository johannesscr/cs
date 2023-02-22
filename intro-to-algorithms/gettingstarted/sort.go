package gettingstarted

import (
	"math"
)

// InsertionSort is a basic sorting algorithm that start from the left of a
// slice of integers, and sequentially moves to the right through all the
// numbers. It sorts by comparing the number at position j, to all previous
// numbers 0 < i < j-1. It moves the number j, to the first position where
// the number at index i is less that the number at index j, and shifts all
// the numbers in between up by one.
func InsertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	return A
}

// SelectionSort applies to sorting n numbers stored in an array A by first
// finding the smallest element of A and exchanging it with the element A[1].
// Then find the second-smallest element of A and exchange it with A[2].
// Continue in this manner for the first n-1 elements of A.
func SelectionSort(A []int) []int {
	n := len(A)
	for i := 0; i < n; i++ {
		// find the i-th element in ascending order
		minA := 0
		minIndex := -1
		for j := i; j < n; j++ {
			if minIndex == -1 || A[j] < minA {
				// condition 1: if is the first iteration of the loop
				// condition 2: if the element A[j] is smaller than the
				// current minimum
				minA = A[j]
				minIndex = j
			}
		}
		// now minA is the current minimum for i-th element in ascending order
		A[i], A[minIndex] = A[minIndex], A[i]
	}
	return A
}

// Merge used to aggregate the two subsequences A[p:q] and A[q+1:r].
// Given A is a slice/array with length >= r, then 0 < p <= q < r are the
// indices of A. The procedure assumes that the sub-arrays , A[p, q] and
// A[q+1, r] are in sorted order.
func Merge(A []int, p, q, r int) []int {
	//fmt.Println("MB", A[p:r+1])
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1+1)  // plus 1 for sentinel element of +infinity
	right := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		left[i] = A[p + i]
	}
	for j := 0; j < n2; j++ {
		right[j] =A[q + j + 1]
	}
	left[n1] = math.MaxInt
	right[n2] = math.MaxInt
	i := 0
	j := 0
	//fmt.Println("L", left[:len(left)-1], "R", right[:len(right)-1])
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			A[k] = left[i]
			i++
		} else {
			A[k] = right[j]
			j++
		}
	}
	//fmt.Println("MA", A[p:r+1])
	return A
}

// MergeSort algorithm which uses divide-and-conquer to sort the sequence.
func MergeSort(A []int, p, r  int) []int {
	if (r-p) > 0 {
		q := (p + r) / 2
		//fmt.Println("MS", A[p:q+1], A[q+1:r+1])
		A = MergeSort(A, p, q)
		A = MergeSort(A, q + 1, r)
		return Merge(A, p, q, r)
	}
	return A
}