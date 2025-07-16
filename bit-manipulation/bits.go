package main

import "fmt"

func main() {
	var x = int8(-5)
	var y = uint8(x)
	// Go just automatically removes any leading 0s.
	fmt.Printf("%b\n", 5)
	fmt.Printf("%b\n", y)
}
