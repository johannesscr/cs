package challenges

// RecurringInt is a Google Question:
//
// The function should take an input which is an array of integers, it should
// return the first recurring integer.
//
// Given an array = [2, 5, 1, 2, 3, 5, 1, 2, 4]
// It should return 2
//
// Given an array = [2, 1, 1, 2, 3, 5, 1, 2, 4]
// It should return 1
//
// Given an array = [2, 3, 4, 5]
// It should return undefined
func RecurringInt(xi []int) (int, bool) {
	hashSet := map[int]bool{}
	for _, i := range xi {
		// check if integer already exists in hash set
		if _, ok := hashSet[i]; ok {
			return i, true
		}
		// add integer to hash set
		hashSet[i] = true
	}
	return 0, false
}
