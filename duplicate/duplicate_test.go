package duplicate

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "no duplicates",
			input:  []int{1, 5, 6, 2, 3, 8},
			output: []int{1, 5, 6, 2, 3, 8},
		},
		{
			name:   "with duplicates",
			input:  []int{1, 5, 2, 6, 3, 2, 7, 8, 9},
			output: []int{1, 5, 2, 6, 3, 7, 8, 9},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := RemoveDuplicates(tc.input)
			out1 := fmt.Sprintf("%v", o)
			out2 := fmt.Sprintf("%v", tc.output)
			if out1 != out2 {
				t.Errorf("expected %v got %v", out2, out1)
			}
		})
	}
}
