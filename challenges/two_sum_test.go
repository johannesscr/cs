package challenges

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			o := TwoSum(tc.nums, tc.target)
			s1 := fmt.Sprintf("%v", tc.output)
			s2 := fmt.Sprintf("%v", o)
			if s1 != s2 {
				t.Errorf("expected output %v got %v", tc.output, o)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	for i := 0; i < b.N; i++ {
		TwoSum(nums, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	for i := 0; i < b.N; i++ {
		TwoSum2(nums, 9)
	}
}
