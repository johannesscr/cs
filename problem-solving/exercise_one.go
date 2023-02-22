package problem_solving

// Given 2 arrays, create a function that lets a user know (true/false) whether
// these two arrays contain any common items.

// example
// array1 = ['a','b','c','d']
// array2 = ['z','y','i']
// should return false

// example
// array1 = ['a','b','x','d']
// array2 = ['z','x','i']
// should return true

// 2 parameters - arrays - no size limit
// return true/false

// Initial brute force solution
// compare every element sequentially in array1 to array2
// that would be a nested loop resulting in O(n^2)

// Next create a hash map/set of the first array1
// then loop through array2 and check if it already existed in array1

func BruteForce(array1, array2 []string) bool {
	for _, element1 := range array1 {
		for _, element2 := range array2 {
			if element1 == element2 {
				return true
			}
		}
	}
	return false
}

func UsingHashMap(array1, array2 []string) bool {
	hashMap := make(map[string]bool)
	// populate hashMap
	for _, element1 := range array1 {
		hashMap[element1] = true
	}
	for _, element2 := range array2 {
		if _, ok := hashMap[element2]; ok {
			return true
		}
	}
	return false
}
