package challenges

import (
	"fmt"
	"testing"
)

/*
542. 01 Matrix
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.

Example 1:

Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]

Example 2:

Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 104
1 <= m * n <= 104
mat[i][j] is either 0 or 1.
There is at least one 0 in mat.
*/

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		{
			name:   "case 1",
			input:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			output: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			name:   "case 2",
			input:  [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			output: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := updateMatrix(test.input)

			m1 := fmt.Sprintf("%v", o)
			m2 := fmt.Sprintf("%v", test.output)
			if m1 != m2 {
				t.Errorf("expect output %s got %s", test.output, o)
			}
		})
	}
}

func updateMatrix(mat [][]int) [][]int {
	return mat
}
