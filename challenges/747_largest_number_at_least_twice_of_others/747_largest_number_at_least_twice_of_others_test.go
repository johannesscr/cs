package _47_largest_number_at_least_twice_of_others

import (
	"testing"
)

/*
747. Largest Number At Least Twice of Others
Easy

You are given an integer array nums where the largest integer is unique.

Determine whether the largest element in the array is at least twice as much as
every other number in the array. If it is, return the index of the largest
element, or return -1 otherwise.

Example 1:

Input: nums = [3,6,1,0]
Output: 1
Explanation: 6 is the largest integer.
For every other number in the array x, 6 is at least twice as big as x.
The index of value 6 is 1, so we return 1.

Example 2:

Input: nums = [1,2,3,4]
Output: -1
Explanation: 4 is less than twice the value of 3, so we return -1.

Constraints:

2 <= nums.length <= 50
0 <= nums[i] <= 100
The largest element in nums is unique.
*/

func TestDominantIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		out  int
	}{
		{
			name: "one",
			nums: []int{3, 6, 1, 0},
			out:  1,
		},
		{
			name: "two",
			nums: []int{1, 2, 3, 4},
			out:  -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := dominantIndex(test.nums)
			if o != test.out {
				t.Errorf("test %s failed, want %d, got %d", test.name, test.out, o)
			}
		})
	}
}

/*
APPROACH ONE
Loop one: find the index of the largest unique number
Loop two: double each value, except the maxVal. keep record of that double max
value, then do the comparison to see if the max unique is still larger than the
double value.
*/

func dominantIndex(nums []int) int {
	index := 0
	maxVal := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if i != index {
			// If the number in the array doubled is larger that the unique max
			// then we already know we can exit and return -1.
			if nums[i]*2 > maxVal {
				return -1
			}
		}
	}

	return index
}
