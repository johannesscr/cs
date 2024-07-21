package codechallenges

import (
	"sort"
)

// StringTrunc takes a string an K the number of characters allowed.
// The function should return the string containing whole words closest to
// the specified number of characters.
//
// messages will not start or end with a space
func StringTrunc(message string, K int) string {
	l := len(message)
	if K >= l {
		return message
	}
	// now K < l
	// now check for the ASCII character 32 which is a space ' '.
	// since the message starts at index 0, checking character k, is like
	// checking the next character in the string
	k := K
	for { // order log n
		if k == 0 {
			break
		}
		if message[k] == 32 {
			break
		}
		k = k - 1
	}
	return message[:k]
}

func sumInts(x ...int) int {
	s := 0
	for _, v := range x {
		s = s + v
	}
	return s
}

// MinCarAllocation finds the minimum number of cars to with which to
// transport all the people.
func MinCarAllocation(p, s []int) int {
	l := len(s)
	totalPeople := sumInts(p...)
	sort.Ints(s)  // sort form
	n := 0
	for i := l-1; i >= 0; i-- {
		if totalPeople <= 0 {
			break
		}
		n++
		totalPeople = totalPeople - s[i]
	}
	return n
}

func convertFloat64(x []int) []float64 {
	y := make([]float64, len(x))
	for i, v := range x {
		y[i] = float64(v)
	}
	return y
}

func sumFloat64s(x ...float64) float64 {
	s := float64(0)
	for i := 0; i < len(x); i++ {
		s = s + x[i]
	}
	return s
}

// PollutionFilters determines the minimum number of filters to half pollution
// the created by the factories fi.
func PollutionFilters(fi []int) int {
	l := len(fi)
	f := convertFloat64(fi) // order n
	goalIndex := sumFloat64s(f...)
	goal := sumFloat64s(f...) / 2
	n := 0
	for goalIndex > goal {
		sort.Float64s(f) // order n log n
		f[l-1] = f[l-1] / 2
		n = n + 1
		goalIndex = sumFloat64s(f...)
	}
	return n
}

// PollutionFilters2 determines the minimum number of filters to half pollution
// the created by the factories fi.
func PollutionFilters2(fi []int) int {
	l := len(fi)
	f := convertFloat64(fi) // order n
	cost := sumFloat64s(f...)
	goal := cost / 2
	n := 0
	for cost > goal {
		sort.Float64s(f) // order n log n
		f[l-1] = f[l-1] / 2
		n = n + 1
		cost = sumFloat64s(f...)
	}
	return n
}

func findMax(x []float64) int {
	m := float64(0)
	idx := 0
	for i, v := range x {
		if v > m {
			m = v
			idx = i
		}
	}
	return idx
}

// PollutionFilters3 determines the minimum number of filters to half pollution
// the created by the factories fi.
func PollutionFilters3(fi []int) int {
	f := convertFloat64(fi) // O(n)
	cost := sumFloat64s(f...)
	goal := cost / 2
	n := 0
	for cost > goal { // O(log n)
		i := findMax(f)
		f[i] = f[i] / 2
		n = n + 1
		cost = cost - f[i]  // reduce cost by amount that has halved
	}
	return n
}

// PollutionFilters4 determines the minimum number of filters to half pollution
// the created by the factories fi.
func PollutionFilters4(fi []int) int {
	l := len(fi)
	//f := convertFloat64(fi)  // order n
	cost := sumInts(fi...)
	goal := cost / 2
	n := 0
	for cost > goal {
		sort.Ints(fi) // order n log n
		fi[l-1] = fi[l-1] / 2
		n = n + 1
		cost = sumInts(fi...)
	}
	return n
}
