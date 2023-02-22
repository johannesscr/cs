package big_o

// Add takes two integers, adds them together and returns the result.
func Add(a, b int) int {
	return a + b
}

// Sum takes a number of integers, adds them all together and returns the
// result.
func Sum(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
