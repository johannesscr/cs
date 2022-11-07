package code_challenge

/*
Given an array of integers nums and an integer target, return indices of
the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you
may not use the same element twice.

You can return the answer in any order.
*/

func TwoSum(nums []int, target int) []int {
	// create a hashmap of compliment values
	// step through the array and add the compliment and the current index
	// to the hashmap
	// nums = 2
	// hashmap 9-2 = 7:0 that means the compliment for seven is at index 2
	// now for the next value in the array I can simply check the hashmap
	// for the location of the compliment.
	compliments := make(map[int]int)

	for i, num := range nums {
		// we have found the compliment
		if compIndex, ok := compliments[num]; ok {
			// since we are stepping through the compliment will always be
			// before the current value
			return []int{compIndex, i}
		}
		// we did not find the compliment
		// add the compliment to the hashmap
		compliments[target-num] = i
	}

	return nil
}

func TwoSum2(nums []int, target int) []int {
	compliments := make(map[int]int, len(nums))
	for i, num := range nums {
		comp := target - num
		if compIndex, ok := compliments[num]; ok {
			return []int{compIndex, i}
		}
		compliments[comp] = i
	}
	return nil
}
