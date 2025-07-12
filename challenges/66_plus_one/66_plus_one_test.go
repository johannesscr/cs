package _6_plus_one

import (
	"reflect"
	"testing"
)

/*
66. Plus One
Easy

You are given a large integer represented as an integer array digits, where
each digits[i] is the ith digit of the integer. The digits are ordered from
most significant to least significant in left-to-right order. The large integer
does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].

Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].

Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].

Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.
*/

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		out    []int
	}{
		{
			name:   "one",
			digits: []int{1, 2, 3},
			out:    []int{1, 2, 4},
		},
		{
			name:   "two",
			digits: []int{4, 3, 2, 1},
			out:    []int{4, 3, 2, 2},
		},
		{
			name:   "three",
			digits: []int{9},
			out:    []int{1, 0},
		},
		{
			name:   "four",
			digits: []int{2, 9, 9},
			out:    []int{3, 0, 0},
		},
		{
			name:   "five",
			digits: []int{9, 9},
			out:    []int{1, 0, 0},
		},
		{
			name:   "six",
			digits: []int{9, 0, 9, 9},
			out:    []int{9, 1, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := plusOne(test.digits)
			if !reflect.DeepEqual(o, test.out) {
				t.Errorf("expected %v, but got %v", test.out, o)
			}
		})
	}
}

/*
APPROACH ONE
Similar to how we would approach long addition. Start at the last digit. If the
digit is 9, then we know we need to set it to 0 and move to the digit before it.
Special case, if the index is 0 and the current digit is 9, then we need to
shift the digits array and add a 1 to the beginning.
*/

func plusOne(digits []int) []int {
	run := true
	i := len(digits) - 1

	for run {
		if digits[i] < 9 {
			digits[i]++
			run = false
		} else {
			digits[i] = 0
			if i > 0 {
				// We are not at the start of the length of digits yet.
				i--
			} else {
				// We are at the zero index.
				digits = append([]int{1}, digits...)
				run = false
			}
		}
	}
	return digits
}

/*
APPROACH TWO
Best Memory

It is actually pretty simple. If the index is 9, then we set it to 0 and only
then do we proceed to the next index. The next index will automatically be
increased not because the previous one was 9. But because we need to increment,
the 9 simply indicated that we need to move to the next index.

If we don't break out of the for loop, then we reached the beginning of the
array and will therefore need to shift the array and add a leading 1.
*/

func plusOneII(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0

	}

	return append([]int{1}, digits...)

}
