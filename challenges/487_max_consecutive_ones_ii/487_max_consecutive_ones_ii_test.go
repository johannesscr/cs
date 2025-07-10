package _87_max_consecutive_ones_ii

import (
	"fmt"
	"testing"
)

/*
487. Max Consecutive Ones II
Medium

Given a binary array nums, return the maximum number of consecutive 1's in the
array if you can flip at most one 0.

Example 1:

Input: nums = [1,0,1,1,0]
Output: 4
Explanation:
- If we flip the first zero, nums becomes [1,1,1,1,0] and we have 4 consecutive ones.
- If we flip the second zero, nums becomes [1,0,1,1,1] and we have 3 consecutive ones.
The max number of consecutive ones is 4.

Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 4
Explanation:
- If we flip the first zero, nums becomes [1,1,1,1,0,1] and we have 4 consecutive ones.
- If we flip the second zero, nums becomes [1,0,1,1,1,1] and we have 4 consecutive ones.
The max number of consecutive ones is 4.

Constraints:

- 1 <= nums.length <= 105
- nums[i] is either 0 or 1.

Follow up:

What if the input numbers come in one by one as an infinite stream? In other
words, you can't store all numbers coming from the stream as it's too large to
hold in memory. Could you solve it efficiently?
*/

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		out  int
	}{
		{
			name: "test one",
			nums: []int{1, 0, 1, 1, 0},
			out:  4,
		},
		{
			name: "test two",
			nums: []int{1, 0, 1, 1, 0, 1},
			out:  4,
		},
		{
			name: "test three",
			nums: []int{1, 1, 1, 1},
			out:  4,
		},
		{
			name: "test four",
			nums: []int{0, 0, 0, 0},
			out:  1,
		},
		{
			name: "test five",
			nums: []int{1, 1, 0, 1},
			out:  4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := findMaxConsecutiveOnes(tc.nums)
			if o != tc.out {
				t.Errorf("expected %d, got %d", tc.out, o)
			}
		})
	}
}

/*
APPROACH ONE:
Keep it simple.

Loop to find all the 0's indices.
Sequentially change the 0 to a 1.
Now we need to loop again, but we can simply from the previous 0 index to the
next 0 index.
*/

func findMaxConsecutiveOnes(nums []int) int {
	xzero := make([]int, 0)
	delta := 0

	// Step through to find all the indices where there are 0s in the array.
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			xzero = append(xzero, i)
		}
	}

	if len(xzero) == 0 {
		return len(nums)
	}

	for i := 0; i < len(xzero); i++ {
		nums[xzero[i]] = 1

		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == 1 {
				count++
			}
			if nums[j] == 0 {
				count = 0
			}
			if count > delta {
				delta = count
			}
		}

		nums[xzero[i]] = 0
	}
	return delta
}

func findMaxConsecutiveOnesI(nums []int) int {
	xzero := make([]int, 0)
	delta := 0

	// Step through to find all the indices where there are 0s in the array.
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			xzero = append(xzero, i)
		}
	}

	//[1, 4]
	// lower = 1
	//i = 4
	// upper = max len

	for i := 0; i < len(xzero); i++ {
		lower := 0
		upper := 0

		if i != 0 {
			lower = xzero[i-1]
		}
		if i == len(xzero)-1 {
			upper = xzero[i]
		} else {
			upper = xzero[i+1]
		}
		di := upper - lower
		fmt.Println(i, lower, upper, di)
		if di > delta {
			delta = di
		}
	}

	return delta
}

/*
FASTEST

This makes use of the two pointer solution.
*/

func findMaxConsecutiveOnesII(nums []int) int {
	maximumNumberOfOnes := 0
	start, end := 0, 0
	lastZero := -1

	for end < len(nums) {

		if nums[end] == 0 {
			if lastZero != -1 {
				start = lastZero + 1
			}
			lastZero = end
		}

		// Increment to next step.
		end++
		// The increment before the max calculation is because end is at the
		// current 0 index, incrementing it is as if we have changed it from
		// a 0 to a 1 and therefore included it in the calculation. Therefore,
		// we do not actually have to the in-place operation.
		maximumNumberOfOnes = max(maximumNumberOfOnes, end-start)
	}

	return maximumNumberOfOnes
}
