package main

import "fmt"

var memoization = make(map[int]int)

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	if memo, ok := memoization[n]; ok {
		return memo
	}
	memoization[n] = fib(n-1) + fib(n-2)
	return memoization[n]
}

func main() {
	fmt.Println(memoization)
	n := fib(100)
	fmt.Println(memoization)
	fmt.Println(n)
}
