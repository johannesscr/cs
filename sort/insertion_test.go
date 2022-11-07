package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "already sorted array",
			input:  []int{1, 2, 3, 4, 5, 6},
			output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "mixed array",
			input:  []int{5, 2, 4, 6, 1, 3},
			output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "array in reverse",
			input:  []int{6, 5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := InsertionSort(tc.input)
			for i := 0; i < len(tc.output); i++ {
				if o[i] != tc.output[i] {
					t.Errorf("expected ordered array position %d to have %d got %d", i, tc.output[i], o[i])
				}
			}
		})
	}
}
