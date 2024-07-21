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
				t.Errorf("expect output %s got %s", m2, m1)
			}
		})
	}
}

func findZero(i, j int, mat [][]int) int {
	// each iteration needs to be a step.
	m := len(mat)       // number of rows.
	n := len(mat[0])    // number of columns.
	q := make([]int, 0) // the queue needs to expand as nodes are added.

	return 0
}

// updateMatrix will step through each of the cells in the array. At each cell
// it will implement a bread-first-search to find the number of steps to get to
// the first `0`.
func updateMatrix(mat [][]int) [][]int {
	m := len(mat)                      // number of rows
	n := len(mat[0])                   // number of columns
	distanceMatrix := make([][]int, m) // make the rows
	for i := 0; i < m; i++ {
		distanceMatrix[i] = make([]int, n) // make columns
	}

	for i, row := range mat {
		for j, cell := range row {
			if cell == 0 {
				break // we can skip this cell
			}
			distance := findZero(i, j, mat)
			distanceMatrix[i][j] = distance
		}
	}
	return mat
}
