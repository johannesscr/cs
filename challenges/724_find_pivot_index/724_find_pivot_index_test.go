package _24_find_pivot_index

import (
	"testing"
)

/*
724. Find Pivot Index
Easy

Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the
left of the index is equal to the sum of all the numbers strictly to the index's
right.

If the index is on the left edge of the array, then the left sum is 0 because
there are no elements to the left. This also applies to the right edge of the
array.

Return the leftmost pivot index. If no such index exists, return -1.

Example 1:

Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11

Example 2:

Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.

Example 3:

Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0

Constraints:

- 1 <= nums.length <= 104
- -1000 <= nums[i] <= 1000

Note:
This question is the same as 1991: https://leetcode.com/problems/find-the-middle-index-in-array/
*/

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		out  int
	}{
		{
			name: "one",
			nums: []int{1, 7, 3, 6, 5, 6},
			out:  3,
		},
		{
			name: "two",
			nums: []int{1, 2, 3},
			out:  -1,
		},
		{
			name: "three",
			nums: []int{2, 1, -1},
			out:  0,
		},
		{
			name: "four",
			nums: []int{-1, -1, -1, -1, -1, 0},
			out:  2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := pivotIndex(test.nums)
			if o != test.out {
				t.Errorf("test %s failed, expect %d but get %d", test.name, test.out, o)
			}
		})
	}
}

func pivotIndex(nums []int) int {
	pivot := len(nums) / 2
	left := 0
	right := 0

	for i := 0; i < pivot; i++ {
		if i < pivot {
			left += nums[i]
		}
		if len(nums)-i-1 > pivot {
			right += nums[len(nums)-i-1]
		}
	}

	for pivot > -1 && pivot < len(nums)-1 {
		if left == right {
			return pivot
		}

		if left > right {
			left -= nums[pivot]
			right += nums[pivot]
			// Shift the pivot to the left
			pivot--
		} else {
			left -= nums[pivot]
			right -= nums[pivot]
			// Shift the pivot to the right
			pivot++
		}
		if pivot == 0 {
			left = 0
		}
		if pivot == len(nums)-1 {
			right = 0
		}
	}

	return -1
}
