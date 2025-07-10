package _48_find_all_numbers_disappeared_in_an_array

import (
	"fmt"
	"sort"
	"testing"
)

/*
448. Find All Numbers Disappeared in an Array
Easy
Given an array nums of n integers where nums[i] is in the range [1, n], return
an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Example 2:

Input: nums = [1,1]
Output: [2]

Constraints:

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n

Follow up:
Could you do it without extra space and in O(n) runtime? You may assume the
returned list does not count as extra space.
*/

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			name: "example 2",
			nums: []int{1, 1},
			want: []int{2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findDisappearedNumbers(tc.nums)
			s1 := fmt.Sprintf("%v", got)
			s2 := fmt.Sprintf("%v", tc.want)
			if s1 != s2 {
				t.Errorf("findDisappearedNumbers(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

/*
APPROACH FOUR
Fastest

Because the nums will only contain numbers [1,n], step through use each nums[i]
as the index to mark the index as being visited by making the value negative.

Finally, step through the array, each nums[i] that is greater than 0 means that
index was not visited, which means the number is missing from the nums array.
*/

func findDisappearedNumbers(nums []int) []int {
	// without using extra space
	// mark the visited index with negative sign
	// visit every element of array, use abs to get correct index and mark it
	result := make([]int, 0)
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	for index, num := range nums {
		if num > 0 {
			result = append(result, index+1)
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

/*
APPROACH THREE
FINAL SOLUTION

Use a hash table to store all the unique integers in nums.
Then step through [1, n] and find all missing number not in the hash table.
*/

func findDisappearedNumbersIII(nums []int) []int {
	hash := make(map[int]bool)
	missing := make([]int, 0)

	for _, num := range nums {
		hash[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := hash[i]; !ok {
			missing = append(missing, i)
		}
	}

	return missing
}

/*
APPROACH TWO

Time limit exceeded.
*/

func findDisappearedNumbersII(nums []int) []int {
	missingNums := make([]int, 0)

	for i := 1; i <= len(nums); i++ {
		exists := false
		for _, num := range nums {
			if num == i {
				exists = true
			}
		}
		if !exists {
			missingNums = append(missingNums, i)
		}
	}
	return missingNums
}

/*
APPROACH ONE
Oops, we misread the question.

Sort the input array nums.
- This will allow us to know the smallest and largest number.
	- The difference between the nums length and the delta difference should
      indicate the size of the missing number array.
- Step through the array.
	- If there is a difference, start at the lower and add the numbers up to
      using the difference as the inner counter.
*/

func findDisappearedNumbersI(nums []int) []int {
	// This is an in-place sort operation
	sort.Ints(nums)
	missingNums := make([]int, 0)

	for i := 1; i < len(nums); i++ {
		delta := nums[i] - nums[i-1] - 1
		if delta > 1 {
			for j := 1; j < delta; j++ {
				missingNums = append(missingNums, nums[i-1]+j)
			}
		}
	}

	return nums
}
