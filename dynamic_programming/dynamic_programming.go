package dynamic_programming

import (
	"fmt"
	"time"
)

func addTo80(n int) int {
	fmt.Println("expensive calculation that takes 2 seconds")
	time.Sleep(2 * time.Second)
	return n + 80
}

// NoMemo does the calculation each and every time the function is called.
func NoMemo(n int) int {
	return addTo80(n)
}

var cache = make(map[int]int)

// Memo uses memoization to speed up calculation if it has already seen
// a certain input.
func Memo(n int) int {
	if v, ok := cache[n]; ok {
		fmt.Println("return value from cache")
		return v
	}
	cache[n] = addTo80(n)
	return cache[n]
}

// Memo2 uses memoization to speed up calculation if it has already seen
// a certain input. Uses a closure to be able to keep the cache within the
// scope of the function, without needing to declare global scope.
func Memo2() func(n int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if v, ok := cache[n]; ok {
			fmt.Println("return value from cache")
			return v
		}
		v := addTo80(n)
		cache[n] = v
		return v
	}
}

var calculationsMemo = 0

func memoFib(index int, cache map[int]int) int {
	if v, ok := cache[index]; ok {
		return v
	}
	calculationsMemo++
	if index < 2 {
		return index
	}
	v := memoFib(index-1, cache) + memoFib(index-2, cache)
	cache[index] = v
	return v
}

func Fib(index int) int {
	cache := make(map[int]int)
	v := memoFib(index, cache)
	fmt.Println(calculationsMemo)
	return v
}

var calculations = 0

func FibNormal(index int) int {
	calculations++
	if index < 2 {
		return index
	}
	return FibNormal(index-1) + FibNormal(index-2)
}

func WrapFibNormal(index int) int {
	v := FibNormal(index)
	fmt.Println(calculations)
	return v
}
