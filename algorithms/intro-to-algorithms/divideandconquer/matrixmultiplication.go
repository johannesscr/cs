package divideandconquer

// MakeMatrix makes a zero matrix of size nxn (n by n).
func MakeMatrix(n int) [][]int {
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, n)
	}
	return A
}

// AddMatrix adds two matrices and takes time O(n^2)
func AddMatrix(A, B [][]int) [][]int {
	n := len(A)
	C := MakeMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

// PartitionMatrix partitions a matrix into 4 quadrants. It assumes that the
// matrix is a multiple of base 2. That is the size of the matrix is
// 2 to the power of x, or n = 2^x.
// example.
// n = 2
// n = 4
// n = 8
// n = 16
func PartitionMatrix(A [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(A)
	P := [][][]int{
		MakeMatrix(n / 2),
		MakeMatrix(n / 2),
		MakeMatrix(n / 2),
		MakeMatrix(n / 2),
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i < n/2 && j < n/2:
				// quadrant 1
				P[0][i][j] = A[i][j]
			case i < n/2 && j >= n/2:
				// quadrant 2
				P[1][i][j-n/2] = A[i][j]
			case i >= n/2 && j < n/2:
				// quadrant 3
				P[2][i-n/2][j] = A[i][j]
			case i >= n/2 && j >= n/2:
				// quadrant 4
				P[3][i-n/2][j-n/2] = A[i][j]
			}
		}
	}
	return P[0], P[1], P[2], P[3]
}

// CombineMatrixPartition takes the four quadrant matrices and combines them
// into a single matrix.
func CombineMatrixPartition(A11, A12, A21, A22 [][]int) [][]int {
	n := 2 * len(A11)
	A := MakeMatrix(n)
	for i := 0; i < n/2; i++ {
		A[i] = append(A11[i], A12[i]...)
		A[i + n/2] = append(A21[i], A22[i]...)
	}
	return A
}

// SquareMatrixMultiplication uses the brute force mathematical definition
// to multiply two square matrices together.
func SquareMatrixMultiplication(A, B [][]int) [][]int {
	n := len(A)
	C := MakeMatrix(n)
	for i := 0; i < n; i++ { // loop through rows i
		for j := 0; j < n; j++ { // loop through columns j
			for k := 0; k < n; k++ { // sum a_{ik} with b_{kj}
				C[i][j] = C[i][j] + (A[i][k] * B[k][j])
			}
		}
	}
	return C
}

// SquareMatrixMultiplyRecursive uses recursion. It assumes that square
// matrices are exact powers of 2 (i.e. divisible by 2 using recursion).
//
// or A = [
// 		[A11 A12]
//		[A21 A22]
// ] and the same for B, then we can calculate C as
// C11 = A11*B11 + A12*B21
// C12 = A11*B12 + A12*B22
// C21 = A21*B11 + A22*B21
// C22 = A21*B12 + A22*B22
func SquareMatrixMultiplyRecursive(A, B [][]int, n int) [][]int {
	if n == 1 {
		return [][]int{
			{A[0][0] * B[0][0]},
		}
	} else {
		// partition A, B and C into 4 quadrants
		A11, A12, A21, A22 := PartitionMatrix(A)
		B11, B12, B21, B22 := PartitionMatrix(B)

		C11 := AddMatrix(SquareMatrixMultiplyRecursive(A11, B11, n/2), SquareMatrixMultiplyRecursive(A12, B21, n/2))
		C12 := AddMatrix(SquareMatrixMultiplyRecursive(A11, B12, n/2), SquareMatrixMultiplyRecursive(A12, B22, n/2))
		C21 := AddMatrix(SquareMatrixMultiplyRecursive(A21, B11, n/2), SquareMatrixMultiplyRecursive(A22, B21, n/2))
		C22 := AddMatrix(SquareMatrixMultiplyRecursive(A21, B12, n/2), SquareMatrixMultiplyRecursive(A22, B22, n/2))
		return CombineMatrixPartition(C11, C12, C21, C22)
	}
}
