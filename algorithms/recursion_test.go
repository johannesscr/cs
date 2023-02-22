package algorithms

import "testing"

func TestFactorialRecursive(t *testing.T) {
	xtc := []struct {
		name   string
		n      int
		output int
	}{
		{
			name:   "2",
			n:      2,
			output: 2,
		},
		{
			name:   "3",
			n:      3,
			output: 6,
		},
		{
			name:   "4",
			n:      4,
			output: 24,
		},
		{
			name:   "5",
			n:      5,
			output: 120,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			o := FactorialRecursive(tc.n)
			if o != tc.output {
				t.Errorf("expected output %d got %d", tc.output, o)
			}
		})
	}
}

func TestFactorialIterative(t *testing.T) {
	xtc := []struct {
		name   string
		n      int
		output int
	}{
		{
			name:   "2",
			n:      2,
			output: 2,
		},
		{
			name:   "3",
			n:      3,
			output: 6,
		},
		{
			name:   "4",
			n:      4,
			output: 24,
		},
		{
			name:   "5",
			n:      5,
			output: 120,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			o := FactorialIterative(tc.n)
			if o != tc.output {
				t.Errorf("expected output %d got %d", tc.output, o)
			}
		})
	}
}

func TestFibonacciRecursive(t *testing.T) {
	xtc := []struct {
		name   string
		index  int
		output int
	}{
		{
			name:   "0",
			index:  0,
			output: 0,
		},
		{
			name:   "1",
			index:  1,
			output: 1,
		},
		{
			name:   "2",
			index:  2,
			output: 1,
		},
		{
			name:   "3",
			index:  3,
			output: 2,
		},
		{
			name:   "4",
			index:  4,
			output: 3,
		},
		{
			name:   "5",
			index:  5,
			output: 5,
		},
		{
			name:   "8",
			index:  8,
			output: 21,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			o := FibonacciRecursive(tc.index)
			if o != tc.output {
				t.Errorf("expected output %d got %d", tc.output, o)
			}
		})
	}
}

func TestFibonacciIterative(t *testing.T) {
	// 0 1 1 2 3 5 8 13 21 34
	xtc := []struct {
		name   string
		index  int
		output int
	}{
		{
			name:   "0",
			index:  0,
			output: 0,
		},
		{
			name:   "1",
			index:  1,
			output: 1,
		},
		{
			name:   "2",
			index:  2,
			output: 1,
		},
		{
			name:   "3",
			index:  3,
			output: 2,
		},
		{
			name:   "4",
			index:  4,
			output: 3,
		},
		{
			name:   "5",
			index:  5,
			output: 5,
		},
		{
			name:   "8",
			index:  8,
			output: 21,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			o := FibonacciIterative(tc.index)
			if o != tc.output {
				t.Errorf("expected output %d got %d", tc.output, o)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	in := "gophers are legit"
	out := "tigel era srehpog"
	o := Reverse(in)
	if o != out {
		t.Errorf("expected reversed string '%s' got '%s'", out, o)
	}
}
