package _24_find_pivot_index

import (
	"fmt"
	"runtime"
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
		{
			name: "five",
			nums: []int{-1, -1, -1, 0, -1, -1},
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

func pivotIndexI(nums []int) int {
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

		// If we have reach the end of the array, set the left or right sum
		// to 0.
		if pivot == 0 {
			left = 0
		}
		if pivot == len(nums)-1 {
			right = 0
		}
	}

	return -1
}

/*
APPROACH TWO
Brute Force

Step
*/

func pivotIndexII(nums []int) int {
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

	basePivot := pivot
	baseLeft := left
	baseRight := right
	direction := 1

	for pivot > -1 && pivot < len(nums)-1 {
		//fmt.Printf("pivot: %d left: %d right: %d\t%v\tnum: %d\n", pivot, left, right, nums, nums[pivot])
		if left == right {
			return pivot
		}

		// Only change direction after we have not found a solution to the left.
		if pivot == 0 && direction == 1 {
			// We were unable to find a solution to the left.
			direction = -1
			// Start again from the initial pivot and search to the right.
			pivot = basePivot
			left = baseLeft
			right = baseRight
		}

		if direction == 1 {
			left -= nums[pivot]
			right += nums[pivot+1]
			// Shift the pivot to the left
			pivot--
		} else {
			left += nums[pivot]
			right -= nums[pivot+1]
			// Shift the pivot to the right
			pivot++
		}
		fmt.Printf("pivot: %d left: %d right: %d\t%v\tnum: %d\n", pivot, left, right, nums, nums[pivot])

		// If we have reach the end of the array, set the left or right sum
		// to 0.
		if pivot == 0 {
			left = 0
		}
		if pivot == len(nums)-1 {
			right = 0
		}
	}

	return -1
}

func calc(nums []int, pivot int) (left, right int) {
	for i := 0; i < len(nums); i++ {
		if i < pivot {
			left += nums[i]
		}
		if pivot < len(nums)-i-1 {
			right += nums[len(nums)-i-1]
		}
	}
	return left, right
}

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left, right := calc(nums, i)
		if left == right {
			return i
		}
	}

	return -1
}

/*
APPROACH FOUR
Most memory effiecient
*/

func pivotIndexIIII(nums []int) int {
	// Strategy
	// Option 1:
	// Brute force this with two for loops

	// Option 2:
	// Try some dynamic programming approach
	// Pre calculate:
	// Input: [1,7,3,6,5,6]
	// DP: [1,8,11,17,22,28]
	// e.g. index 0
	// left 0
	// dp[len-1]-nums[i] = 26
	// e.g. index 3
	// left dp[i-1] = 11
	// right dp[len-1]=28 - dp[i]=17 = 11

	runtime.GC()

	var totalSum int
	for i := range nums {
		totalSum += nums[i]
	}

	currSum := 0
	for i := 0; i < len(nums); i++ {
		if currSum == totalSum-currSum-nums[i] {
			return i
		}

		currSum += nums[i]
	}

	return -1
}

/*
APPROACH FIVE
Fastest
*/
func pivotIndexIIIII(nums []int) int {
	left := 0
	right := 0

	for _, num := range nums {
		right += num
	}

	for i, num := range nums {
		// Subtract the num (which is the pivot) from the right total so that
		// we do not include the pivot in the right sum.
		if left == right-num {
			return i
		}
		// Add the num to the left and subtract from the right, similar to
		// moving the pivot and recalculating.
		left += num
		right -= num
	}

	return -1
}
