package algorithms

var counter = 0

func Recur() string {
	if counter > 3 {
		return "done"
	}
	counter++
	return Recur()
}

func FactorialRecursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

func FactorialIterative(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total = total * i
	}
	return total
}

// FibonacciRecursive is O(2^n) exponential time
func FibonacciRecursive(index int) int {
	if index < 2 {
		// that is 0 or 1
		return index
	}
	return FibonacciRecursive(index-1) + FibonacciRecursive(index-2)
}

// FibonacciIterative is O(n) linear time
func FibonacciIterative(index int) int {
	if index < 2 {
		// either 0 or 1
		return index
	}
	previous := []int{0, 1}
	for i := 2; i < index; i++ {
		pm0 := previous[0]
		pm1 := previous[1]
		previous[1] = pm1 + pm0
		previous[0] = pm1
	}
	return previous[0] + previous[1]
}

func Reverse(s string) string {
	if len(s) == 1 {
		return s
	}
	xb := []byte(s)
	lastElement := string(xb[len(xb)-1])
	remainingString := string(xb[:len(xb)-1])
	return lastElement + Reverse(remainingString)
}
