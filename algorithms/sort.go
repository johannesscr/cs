package algorithms

func BubbleSort(xi []int) []int {
	for k := 0; k < len(xi); k++ {
		for i := 0; i < len(xi)-1-k; i++ {
			if xi[i] > xi[i+1] {
				// can do a dual swap in go
				xi[i], xi[i+1] = xi[i+1], xi[i]
			}
		}
	}
	return xi
}

func SelectionSort(xi []int) []int {
	for i := 0; i < len(xi); i++ {
		// i, denotes the position to be swapped into in the slice.
		// that is also why we start looping from i.
		m := xi[i] // current minimum.
		index := i
		for j := i + 1; j < len(xi); j++ {
			if xi[j] < m {
				// element is smaller than the current minimum.
				m = xi[j]
				index = j
			}
		}
		xi[i], xi[index] = xi[index], xi[i]
	}
	return xi
}

func InsertionSort(xi []int) []int {
	newXI := []int{xi[0]}
	for i := 1; i < len(xi); i++ {
		v := xi[i]
		if v <= newXI[0] {
			newXI = append([]int{v}, newXI...)
		} else if v >= newXI[len(newXI)-1] {
			newXI = append(newXI, v)
		} else {
			for j := 1; j < i; j++ {
				//fmt.Println(i, j, ":", v, newXI[j], ":", len(newXI), newXI)
				if v < newXI[j] {
					start := newXI[:j]
					end := append([]int{v}, newXI[j:]...)
					newXI = append(start, end...)
					break
				}
			}
		}
	}
	return newXI
}

func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	// split the array
	left, right := split(a)

	return merge2(MergeSort(left), MergeSort(right))
}

func split(a []int) ([]int, []int) {
	left := a[:len(a)/2]
	right := a[len(a)/2:]
	return left, right
}

func merge(a, b []int) []int {
	sorted := make([]int, 0)
	length := len(a) + len(b)
	for len(sorted) < length {
		switch {
		case len(a) == 0:
			sorted = append(sorted, b...)
		case len(b) == 0:
			sorted = append(sorted, a...)
		default:
			if a[0] < b[0] {
				sorted = append(sorted, a[0])
				a = a[1:]
			} else {
				sorted = append(sorted, b[0])
				b = b[1:]
			}
		}
	}
	return sorted
}

func merge2(a, b []int) []int {
	sorted := make([]int, 0)
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(a) && rightIndex < len(b) {
		if a[leftIndex] < b[rightIndex] {
			sorted = append(sorted, a[leftIndex])
			leftIndex++
		} else {
			sorted = append(sorted, b[rightIndex])
			rightIndex++
		}
	}
	remain := append(a[leftIndex:], b[rightIndex:]...)
	return append(sorted, remain...)
}

func QuickSort() []int {
	// still to implement
	return nil
}
