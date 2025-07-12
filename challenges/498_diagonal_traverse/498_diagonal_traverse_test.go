package _98_diagonal_traverse

import (
	"reflect"
	"testing"
)

/*
498. Diagonal Traverse
Medium

Given an m x n matrix mat, return an array of all the elements of the array in
a diagonal order.

Example 1:

Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]

Example 2:

Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 104
1 <= m * n <= 104
-105 <= mat[i][j] <= 105
*/

func TestFindDiagonalOrder(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		out  []int
	}{
		{
			name: "one",
			mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			out:  []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "two",
			mat:  [][]int{{1, 2}, {3, 4}},
			out:  []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := findDiagonalOrder(test.mat)
			if !reflect.DeepEqual(o, test.out) {
				t.Errorf("expected %v, but got %v", test.out, o)
			}
		})
	}
}

/*
APPROACH ONE
Let's look at the i, j indices of the matrix

[1,2,3]
[4,5,6]
[7,8,9]

We have the numbers as
1  2  4  7  5  3  6  8  9
11 12 21 31 22 13 23 32 33
up dn dn up up up dn dn up

Notice that we have groups of
2  3  3  4  4  4  5  5  6
Where the max of 6 is the sum of N+M (3 + 3).

The order changes:
i small and j large when going up and
i large and j small when going down.

Now wat we do is iterate over the groups [2,M+N], next we toggle the direction
up and down for each group. Within the group up as increment and do the
complementary calculation such as j = x - i. For this we need a non-square
matrix.

[1,2]
[3,4]
[5,6]

1  2  3  5  4  6
11 12 21 31 22 32
2  3  3  4  4  5
up dn dn up up dn

On the up we increment j and find i as the compliment.
On the down we increment i and find j as the compliment.
*/

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	flat := make([]int, m*n)
	up := true
	index := 0

	for g := 2; g <= m+n; g++ {
		//fmt.Printf("group: %d up:%t\n", g, up)
		for x := 1; x <= g; x++ {
			if up {
				j := x
				i := g - j
				//fmt.Printf("group: %d i:%d, j:%d ", g, i, j)
				if i > m || i == 0 || j > n {
					// If row index is larger that the total number of rows
					// this is invalid
					//fmt.Printf("skip\n")
					continue
				}
				//fmt.Printf("add \n")
				flat[index] = mat[i-1][j-1]
				index++
			} else {
				i := x
				j := g - i
				if j > n || j == 0 || i > m {
					// If the column index is larger than the total number of
					// columns this is invalid
					//fmt.Printf("group: %d i:%d, j:%d skip\n", g, i, j)
					continue
				}
				//fmt.Printf("group: %d i:%d, j:%d\n", g, i, j)
				flat[index] = mat[i-1][j-1]
				index++
			}
		}
		// Toggle direction.
		up = !up
	}

	return flat
}

/*
APPROACH TWO
Fastest

           +
  [1,2,3]  | 2 rows down
  [4,5,6]<-+ 1 col back
  [7,8,9]
     ^
+----+ 1 row up.
2 cols forward
*/

func findDiagonalOrderII(mat [][]int) []int {
	if len(mat) == 0 {
		return nil
	}

	rows := len(mat)
	cols := len(mat[0])

	results := make([]int, 0, rows+cols)

	goingUp := true

	row, col := 0, 0
	for row != rows && col != cols {
		if goingUp {
			for col < cols && row >= 0 {
				results = append(results, mat[row][col])
				row--
				col++
			}

			// This is the switch to get the starting point for going back down.
			if col < cols {
				row++
			} else {
				// We have reached the top right corner of the matrix. We have
				// over shot the column by 1 now. But our next starting row
				// is two rows down, because we are at the top right corner.
				row += 2
				col--
			}

		} else {
			for row < rows && col >= 0 {
				results = append(results, mat[row][col])
				row++
				col--
			}

			// This is the switch to get the starting point for going back up.
			if row < rows {
				// return column from out of bounds
				col++
			} else {
				row--
				col += 2
			}
		}

		goingUp = !goingUp
	}

	return results
}
