package _14_third_maximum_number

import (
	"math"
	"testing"
)

/*
414. Third Maximum Number
Easy
Given an integer array nums, return the third distinct maximum number in this
array. If the third maximum does not exist, return the maximum number.

Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.

Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.

Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.

Constraints:

- 1 <= nums.length <= 104
- -231 <= nums[i] <= 231 - 1

Follow up:
Can you find an O(n) solution?
*/

func TestThirdMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		out  int
	}{
		{
			name: "one",
			nums: []int{3, 2, 1},
			out:  1,
		},
		{
			name: "two",
			nums: []int{1, 2},
			out:  2,
		},
		{
			name: "three",
			nums: []int{2, 2, 3, 1},
			out:  1,
		},
		{
			name: "four",
			nums: []int{5, 2, 2},
			out:  5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := thirdMax(tc.nums)
			if o != tc.out {
				t.Errorf("thirdMax(%v) = %v, want %v", tc.nums, o, tc.out)
			}
		})
	}
}

/*
APPROACH ONE
Keep three variables. For each distinct max.
Initialize to the first value in the array.
Then step through the array and do three checks, if the value is greater than
any of the maxes, replace the max and continue to the next value.
*/

func thirdMaxI(nums []int) int {
	maxI := math.MinInt
	maxII := math.MinInt
	maxIII := math.MinInt
	for _, num := range nums {
		if num > maxI {
			maxI = num
			continue
		}
		if num > maxII {
			maxII = num
			continue
		}
		if num > maxIII {
			maxIII = num
			continue
		}
	}
	if maxIII > math.MinInt {
		return maxIII
	}
	return maxI
}

/*
APPROACH TWO
Keep an array of length 3, then push onto the array. If the array has length
less than 3, return max, else return the third max.
*/

func thirdMax(nums []int) int {
	m := []int{math.MinInt, math.MinInt, nums[0]}

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		// Now check where to add the number.
		if x > m[2] {
			// Swap the numbers so that we can filter down to the next index.
			m[2], x = x, m[2]
		}
		if x > m[1] && x < m[2] {
			m[1], x = x, m[1]
		}
		if x > m[0] && x < m[1] {
			m[0], x = x, m[0]
		}
	}

	if m[0] == math.MinInt {
		// Then we need to return the maximum in the slice.
		return m[2]
	}

	return m[0]
}
