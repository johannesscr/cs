package divideandconquer

import (
	"math"
)

func FindMaxCrossingSubarray(A []int, low, mid, high int) (int, int, int) {
	leftSum := math.MinInt
	sum := 0
	maxLeft := mid
	for i := mid; i >= low ; i-- {
		sum = sum + A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := math.MinInt
	sum = 0
	maxRight := mid + 1
	for j := mid + 1; j < high; j++ {
		sum = sum + A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	// go -math.MinInt minus a negative value => positive value
	if leftSum < 0 && rightSum == math.MinInt {
		return maxLeft, maxRight, leftSum
	}
	if rightSum < 0 && leftSum == math.MinInt {
		return maxLeft, maxRight, rightSum
	}
	return maxLeft, maxRight, leftSum + rightSum
}

// FindMaxSubarray is a function that uses divide and conquer to find the
// maximum subarray in n lg(n) time or O(n lg n). Recall log is base 2.
func FindMaxSubarray(A []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, A[low]
	} else {
		mid := (low + high) / 2
		// int will automatically truncate to lower-bound int
		leftLow, leftHigh, leftSum := FindMaxSubarray(A, low, mid)
		rightLow, rightHigh, rightSum := FindMaxSubarray(A, mid+1, high)
		crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(A, low, mid, high)

		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

// MaxSubarray should be a function that can find the maximum sub array in
// linear time or O(n)
func MaxSubarray(A []int) (int, int, int) {
	var high int
	low := 0
	sum := 0
	maxSum := 0

	//fmt.Printf("i\tlow\thigh\tA[i]\t\tsum\tmax sum\n")
	for i := 0; i < len(A); i++ {
		// reset
		if sum < 0 && A[i] > sum {
			low = i
			high = i
			sum = A[i]
		} else if sum + A[i] > maxSum {
			// if sum is the new max, then update the max subarray
			high = i

			sum = sum + A[i]
			maxSum = sum
		} else {
			// no update and continue searching
			sum = sum + A[i]
		}
		//fmt.Printf("%d\t%d\t%d\t%d\t>\t%d\t%d\n", i, low, high, A[i], sum, maxSum)
	}
	return low, high, maxSum
}