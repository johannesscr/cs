package big_o

import "testing"

/*
	O(1) CONSTANT TIME
*/

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

/*
	O(n) LINEAR TIME
*/

func BenchmarkSum_1(b *testing.B) {
	ints := make([]int, 1)
	for i := 0; i < b.N; i++ {
		Sum(ints...)
	}
}

func BenchmarkSum_1000(b *testing.B) {
	ints := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		Sum(ints...)
	}
}

func BenchmarkSum_1000000(b *testing.B) {
	ints := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		Sum(ints...)
	}
}
