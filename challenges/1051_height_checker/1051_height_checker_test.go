package _051_height_checker

import (
	"sort"
	"testing"
)

/*
1051. Height Checker
#easy #array

A school is trying to take an annual photo of all the students. The students are
asked to stand in a single file line in non-decreasing order by height. Let this
ordering be represented by the integer array expected where expected[i] is the
expected height of the ith student in line.

You are given an integer array heights representing the current order that the
students are standing in. Each heights[i] is the height of the ith student in
line (0-indexed).

Return the number of indices where heights[i] != expected[i].

Example 1:

Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation:
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.

Example 2:

Input: heights = [5,1,2,3,4]
Output: 5
Explanation:
heights:  [5,1,2,3,4]
expected: [1,2,3,4,5]
All indices do not match.

Example 3:

Input: heights = [1,2,3,4,5]
Output: 0
Explanation:
heights:  [1,2,3,4,5]
expected: [1,2,3,4,5]
All indices match.

Constraints:

1 <= heights.length <= 100
1 <= heights[i] <= 100
*/

func TestHeightChecker(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "example one",
			input:  []int{1, 1, 4, 2, 1, 3},
			output: 3,
		},
		{
			name:   "example two",
			input:  []int{5, 1, 2, 3, 4},
			output: 5,
		},
		{
			name:   "example three",
			input:  []int{1, 2, 3, 4, 5},
			output: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := heightChecker(tc.input)
			if o != tc.output {
				t.Errorf("expected %d, got %d", tc.output, o)
			}
		})
	}
}

/*
APPROACH:
Use merge sort to create a copy of the input array that is sorted. Then step
through the array, at each index increment the diff counter if the entry at that
index differs.
*/
func heightChecker(heights []int) int {
	// here we copy heights
	sorted := make([]int, len(heights))
	copy(sorted, heights)

	sort.Ints(sorted)

	diff := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sorted[i] {
			diff++
		}
	}

	return diff
}

/*
APPROACH I
*/
func heightCheckerI(heights []int) int {
	swapp := true
	ans := 0
	arr := make([]int, len(heights))
	aux := 0

	// copy each element manually
	for i := 0; i < len(heights); i++ {
		arr[i] = heights[i]
	}

	// use bubble sort to sort the copied array
	for swapp {
		swapp = false
		for i := 0; i < len(heights)-1; i++ {
			if heights[i] > heights[i+1] {
				aux = heights[i]
				heights[i] = heights[i+1]
				heights[i+1] = aux
				swapp = true
			}
		}
	}

	// do the comparison as before
	for i := 0; i < len(heights); i++ {
		if arr[i] != heights[i] {
			ans++
		}
	}

	return ans
}
