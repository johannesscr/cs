package data_structure

// In Go, we have two ways of creating an array
// var slice []int   // this has a mutable length as needed
// var array [10]int // this has fixed length of 10
// since there is fixed length of 10, for a 32 bit machine each allocation is
// 4 bytes, therefore for this array a fixed memory is needed of 10*4 = 40 bytes

/* LET'S TRY OUR OWN ARRAY */

//type Array map[int]interface{}

type Array struct {
	data   map[int]interface{}
	length int
}

func NewArray() Array {
	return Array{
		data:   make(map[int]interface{}),
		length: 0,
	}
}

// Get returns the value at the specified index, if the index is out of bounds
// it will return nil.
//
// Note this function is O(1)
func (a *Array) Get(index int) interface{} {
	return a.data[index] // O(1)
}

// Push appends a new element to the end of the array and increments the length
// of the Array.
//
// Note this function is O(1)
func (a *Array) Push(v interface{}) {
	a.data[a.length] = v // O(1)
	a.length += 1        // O(1)
}

// Pop removes the last element from the array.
//
// Note this is O(1) Constant time.
func (a *Array) Pop() interface{} {
	v := a.data[a.length-1]    // O(1)
	delete(a.data, a.length-1) // O(1)
	a.length -= 1              // O(1)
	return v
}

// Delete removes an element at a specified index.
//
// Note that this function is O(n) Linear time.
func (a *Array) Delete(index int) interface{} {
	v := a.data[index]     // O(1)
	a.shiftElements(index) // O(n)
	return v
}

// shiftElements shifts all the elements to the right of the index to the left
// by one, and then deletes the last element.
//
// Note due to the loop this function is O(n) Linear time.
func (a *Array) shiftElements(index int) {
	for i := index; i < a.length-1; i++ { // Loop causes O(n)
		// shift all items left by one
		a.data[i] = a.data[i+1]
	}
	// now delete the last item
	// using the method we've already created
	a.Pop() // O(1)
}

/*
 *******************************************************************************
 SAMPLE QUESTION
 *******************************************************************************
*/

/* QUESTION 1 */

// Given an array of integers, return the array in reverse order

/* QUESTION 2 */

// Write a function that takes two arrays, then return a single array with the
// two arrays combined, and is sorted.
//
// Suppose MergeSortedArrays([0, 3, 4, 31], [4, 6, 30])
// Output: [0, 3, 4, 4, 6, 30, 31]

func MergeSortedArrays(a, b []int) []int {
	// calculate the total length the final array should be
	lenA := len(a)
	lenB := len(b)
	maxA := a[lenA-1]
	maxB := b[lenB-1]
	newArray := make([]int, lenA+lenB)
	counterA := 1
	counterB := 1
	ai := a[0]
	bi := b[0]

	inf := maxA + 1
	if maxB > maxA {
		inf = maxB + 1
	}

	// get a value from each array
	for i := 0; i < lenA+lenB; i++ {
		// get the lowest value and add it to the new array default use the first
		if bi < ai {
			// if the second is smaller use that one
			newArray[i] = bi
			bi = unshiftOrInf(b, &counterB, inf)
		} else {
			// else use the value from the first array
			newArray[i] = ai
			ai = unshiftOrInf(a, &counterA, inf)
		}
	}

	return newArray
}

func unshiftOrInf(array []int, index *int, maxValue int) int {
	if *index == len(array) {
		return maxValue
	}
	*index += 1
	return array[*index-1]
}

// Javascript example does not translate directly to Go.
//func MergeSortedArrayExample(a, b []int) []int {
//	mergedArray := make([]int, 0)
//	ai := a[0]
//	bj := b[0]
//	i := 1
//	j := 1
//
//	// check input, for not typed languages
//	if len(a) == 0 {
//		return b
//	}
//	if len(b) == 0 {
//		return a
//	}
//
//	for j <= len(b) || i <= len(a) {
//		if i == len(a) || ai < bj {
//			mergedArray = append(mergedArray, ai)
//			ai = a[i]
//			i++
//		} else {
//			mergedArray = append(mergedArray, bj)
//			bj = b[j]
//			j++
//		}
//	}
//
//	return mergedArray
//}
