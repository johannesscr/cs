package duplicate

func RemoveDuplicates(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := 1; j < len(a); j++ {
			if a[i] == a[j] {
				a = append(a[:j], a[j+1:]...)
			}
		}
	}
	return a
}
