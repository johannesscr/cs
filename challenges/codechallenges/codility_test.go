package codechallenges

import "testing"

func TestFindMax(t *testing.T) {
	tt := []struct {
		name  string
		xf    []float64
		index int
	}{
		{
			xf:    []float64{0, 12, 12, 0},
			index: 1,
		},
		{
			xf:    []float64{0, 0, 0, 0},
			index: 0,
		},
		{
			xf:    []float64{5, 5, 5, 5},
			index: 0,
		},
		{
			xf:    []float64{1, 2, 3, 4},
			index: 3,
		},
		{
			xf:    []float64{1, 2, 4, 3},
			index: 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			idx := findMax(tc.xf)
			if idx != tc.index {
				t.Errorf("expected index %d got %d", tc.index, idx)
			}
		})
	}
}

func TestStringTrunc(t *testing.T) {
	tt := []struct {
		name    string
		message string
		K       int
		output  string
	}{
		{
			name:    "zero k",
			message: "this is a random message",
			K:       0,
			output:  "",
		},
		{
			name:    "two words",
			message: "this is a random message",
			K:       7,
			output:  "this is",
		},
		{
			name:    "zero words",
			message: "this is a random message",
			K:       2,
			output:  "",
		},
		{
			name:    "five words",
			message: "this is a random message",
			K:       25,
			output:  "this is a random message",
		},
		{
			name:    "three words",
			message: "this is a random message",
			K:       15,
			output:  "this is a",
		},
		{
			name:    "four words",
			message: "this is a random message",
			K:       16,
			output:  "this is a random",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			m := StringTrunc(tc.message, tc.K)
			if m != tc.output {
				t.Errorf("expected '%s' got '%s'", tc.output, m)
			}
		})
	}
}

func BenchmarkStringTrunc(b *testing.B) {
	xb := []struct {
		message string
		K       int
	}{
		{
			message: "this is a random message",
			K:       7,
		},
		{
			message: "this is a random message",
			K:       2,
		},
		{
			message: "this is a random message",
			K:       25,
		},
		{
			message: "this is a random message",
			K:       17,
		},
	}

	for i := 0; i < b.N; i++ {
		m := xb[i%4].message
		k := xb[i%4].K
		StringTrunc(m, k)
	}
}

func TestPollutionFilters(t *testing.T) {
	tt := []struct {
		name        string
		filterIndex []int
		output      int
	}{
		{
			name:        "",
			filterIndex: []int{1, 19},
			output:      2,
		},
		{
			name:        "",
			filterIndex: []int{0, 0, 0, 0},
			output:      0,
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 0},
			output:      1,
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 1},
			output:      2,
		},
		{
			name:        "",
			filterIndex: []int{1, 1, 0, 1},
			output:      3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := PollutionFilters(tc.filterIndex)
			if o != tc.output {
				t.Errorf("expected %d got %d", tc.output, o)
			}
		})
	}
}

func BenchmarkPollutionFilters(b *testing.B) {
	xBench := []struct {
		factories []int
	}{
		{factories: []int{0, 0, 0, 1, 0, 9, 12}},
		{factories: []int{0, 19, 172, 16, 19, 46, 1, 54, 0, 0, 24}},
		{factories: []int{91, 46, 15, 49, 17, 46, 1, 0, 18, 57, 17, 57}},
		{factories: []int{19, 1}},
	}
	for i := 0; i < b.N; i++ {
		PollutionFilters(xBench[i%4].factories)
	}
}

func TestPollutionFilters2(t *testing.T) {
	tt := []struct {
		name        string
		filterIndex []int
		output      int
	}{
		{
			name:        "",
			filterIndex: []int{1, 19},
			output:      2,
		},
		{
			name:        "",
			filterIndex: []int{0, 0, 0, 0},
			output:      0,
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 0},
			output:      1,
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 1},
			output:      2,
		},
		{
			name:        "",
			filterIndex: []int{1, 1, 0, 1},
			output:      3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := PollutionFilters2(tc.filterIndex)
			if o != tc.output {
				t.Errorf("expected %d got %d", tc.output, o)
			}
		})
	}
}

func BenchmarkPollutionFilters2(b *testing.B) {
	xBench := []struct {
		factories []int
	}{
		{factories: []int{0, 0, 0, 1, 0, 9, 12}},
		{factories: []int{0, 19, 172, 16, 19, 46, 1, 54, 0, 0, 24}},
		{factories: []int{91, 46, 15, 49, 17, 46, 1, 0, 18, 57, 17, 57}},
		{factories: []int{19, 1}},
	}
	for i := 0; i < b.N; i++ {
		PollutionFilters2(xBench[i%4].factories)
	}
}

func TestPollutionFilters3(t *testing.T) {
	tt := []struct {
		name        string
		filterIndex []int
		output      int
	}{
		{
			name:        "",
			filterIndex: []int{1, 19},
			output:      2,
		},
		{
			name:        "",
			filterIndex: []int{0, 0, 0, 0},
			output:      0,
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 0},
			output:      1,
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 1},
			output:      2,
		},
		{
			name:        "",
			filterIndex: []int{1, 1, 0, 1},
			output:      3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := PollutionFilters3(tc.filterIndex)
			if o != tc.output {
				t.Errorf("expected %d got %d", tc.output, o)
			}
		})
	}
}

func BenchmarkPollutionFilters3(b *testing.B) {
	xBench := []struct {
		factories []int
	}{
		{factories: []int{0, 0, 0, 1, 0, 9, 12}},
		{factories: []int{0, 19, 172, 16, 19, 46, 1, 54, 0, 0, 24}},
		{factories: []int{91, 46, 15, 49, 17, 46, 1, 0, 18, 57, 17, 57}},
		{factories: []int{19, 1}},
	}
	for i := 0; i < b.N; i++ {
		PollutionFilters3(xBench[i%4].factories)
	}
}

func TestPollutionFilters4(t *testing.T) {
	tt := []struct {
		name        string
		filterIndex []int
		output      int
	}{
		{
			name:        "",
			filterIndex: []int{1, 19},
			output:      1, // rounding error
		},
		{
			name:        "",
			filterIndex: []int{0, 0, 0, 0},
			output:      0,
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 0},
			output:      1, // rounding error
		},
		{
			name:        "",
			filterIndex: []int{0, 1, 0, 1},
			output:      1, // rounding error
		},
		{
			name:        "",
			filterIndex: []int{1, 1, 0, 1},
			output:      2, // rounding error
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := PollutionFilters4(tc.filterIndex)
			if o != tc.output {
				t.Errorf("expected %d got %d", tc.output, o)
			}
		})
	}
}

func BenchmarkPollutionFilters4(b *testing.B) {
	xBench := []struct {
		factories []int
	}{
		{factories: []int{0, 0, 0, 1, 0, 9, 12}},
		{factories: []int{0, 19, 172, 16, 19, 46, 1, 54, 0, 0, 24}},
		{factories: []int{91, 46, 15, 49, 17, 46, 1, 0, 18, 57, 17, 57}},
		{factories: []int{19, 1}},
	}
	for i := 0; i < b.N; i++ {
		PollutionFilters4(xBench[i%4].factories)
	}
}

func TestMinCarAllocation(t *testing.T) {
	tt := []struct {
		name string
		p []int
		s []int
		output int
	}{
		{
			p: []int{},
			s: []int{},
			output: 0,
		},
		{
			p: []int{1, 5, 2},
			s: []int{1, 6, 4},
			output: 2,
		},
		{
			p: []int{1, 5, 2, 1, 1},
			s: []int{1, 6, 4, 2, 2},
			output: 2,
		},
		{
			p: []int{1, 5, 2, 1, 1},
			s: []int{1, 6, 3, 2, 2},
			output: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := MinCarAllocation(tc.p, tc.s)
			if o != tc.output {
				t.Errorf("expected number of cars %d got %d", tc.output, o)
			}
		})
	}
}

func BenchmarkMinCarAllocation(b *testing.B) {
	xBench := []struct {
		p []int
		s []int
	}{
		{
			p: []int{},
			s: []int{},
		},
		{
			p: []int{1, 5, 2},
			s: []int{1, 6, 4},
		},
		{
			p: []int{1, 5, 2, 1, 1},
			s: []int{1, 6, 4, 2, 2},
		},
		{
			p: []int{1, 5, 2, 1, 1},
			s: []int{1, 6, 3, 2, 2},
		},
	}

	for i := 0; i < b.N; i++ {
		MinCarAllocation(xBench[i%4].p, xBench[i%4].s)
	}
}

func BenchmarkMinCarAllocation2(b *testing.B) {
	xBench := []struct {
		p []int
		s []int
	}{
		{
			p: []int{},
			s: []int{},
		},
		{
			p: []int{1, 5, 2, 1, 5, 2},
			s: []int{1, 6, 4, 1, 6, 4},
		},
		{
			p: []int{1, 5, 2, 1, 1, 1, 5, 2, 1, 1},
			s: []int{1, 6, 4, 2, 2, 1, 6, 4, 2, 2},
		},
		{
			p: []int{1, 5, 2, 1, 1, 1, 5, 2, 1, 1},
			s: []int{1, 6, 3, 2, 2, 1, 6, 3, 2, 2},
		},
	}

	for i := 0; i < b.N; i++ {
		MinCarAllocation(xBench[i%4].p, xBench[i%4].s)
	}
}

func BenchmarkMinCarAllocation4(b *testing.B) {
	xBench := []struct {
		p []int
		s []int
	}{
		{
			p: []int{},
			s: []int{},
		},
		{
			p: []int{1, 5, 2, 1, 5, 2, 1, 5, 2, 1, 5, 2},
			s: []int{1, 6, 4, 1, 6, 4, 1, 6, 4, 1, 6, 4},
		},
		{
			p: []int{1, 5, 2, 1, 1, 1, 5, 2, 1, 1, 1, 5, 2, 1, 1, 1, 5, 2, 1, 1},
			s: []int{1, 6, 4, 2, 2, 1, 6, 4, 2, 2, 1, 6, 4, 2, 2, 1, 6, 4, 2, 2},
		},
		{
			p: []int{1, 5, 2, 1, 1, 1, 5, 2, 1, 1, 1, 5, 2, 1, 1, 1, 5, 2, 1, 1},
			s: []int{1, 6, 3, 2, 2, 1, 6, 3, 2, 2, 1, 6, 3, 2, 2, 1, 6, 3, 2, 2},
		},
	}

	for i := 0; i < b.N; i++ {
		MinCarAllocation(xBench[i%4].p, xBench[i%4].s)
	}
}
