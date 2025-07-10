package _77_squares_of_a_sorted_array

import (
	"math"
	"reflect"
	"testing"
)

/*
977. Squares of a Sorted Array
Easy
Given an integer array nums sorted in non-decreasing order, return an array of
the squares of each number sorted in non-decreasing order.

Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Constraints:

- 1 <= nums.length <= 104
- -104 <= nums[i] <= 104
- nums is sorted in non-decreasing order.

Follow up:
Squaring each element and sorting the new array is very trivial, could you find
an O(n) solution using a different approach?
*/

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "example 2",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sortedSquares(test.nums)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

/*
APPROACH ONE
Step through the array nums.
- Square the value at the given index.
- Keep track of the value closet to 0 index.

From the minimum value's index, step both left and right outwards to create
a new array while sorting the square values.
*/

func sortedSquares(nums []int) []int {
	// Minimum Index, Value closest to Zero
	vi, v := 0, 104

	for i := 0; i < len(nums); i++ {
		if abs(nums[i]) < v {
			v = abs(nums[i])
			vi = i
		}
		nums[i] = nums[i] * nums[i]
	}

	// Now we have the array all squared, now we need to sort the array.
	left := vi - 1
	right := vi + 1
	lv := -1
	rv := -1
	si := 1
	sorted := make([]int, len(nums))
	sorted[0] = nums[vi]

	for i := 0; i < len(nums); i++ {
		if left < 0 && right >= len(nums) {
			break
		}

		if left >= 0 {
			lv = nums[left]
		} else {
			lv = math.MaxInt
		}
		if right < len(nums) {
			rv = nums[right]
		} else {
			rv = math.MaxInt
		}

		if lv <= rv {
			sorted[si] = lv
			left--
		} else {
			sorted[si] = rv
			right++
		}
		si++
	}
	return sorted
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
APPROACH TWO
FASTEST

This approach is to work inwards from the ends. This is because when the ends
are squared they will result in the largest numbers. Then simply add the largest
to the end and take one step inwards.
*/

func sortedSquaresI(nums []int) []int {
	result := make([]int, len(nums))

	end := len(nums) - 1
	i, j := 0, end
	for i <= j {
		n1 := nums[i] * nums[i]
		n2 := nums[j] * nums[j]
		if n1 > n2 {
			i++
		} else {
			j--
		}

		result[end] = max(n1, n2)
		end--
	}
	return result
}
