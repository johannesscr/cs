package problem_solving

import "testing"

func TestBruteForce(t *testing.T) {
	tt := []struct {
		name   string
		array1 []string
		array2 []string
		output bool
	}{
		{
			name:   "no match",
			array1: []string{"a", "b", "c", "d"},
			array2: []string{"x", "y", "z"},
			output: false,
		},
		{
			name:   "match",
			array1: []string{"a", "b", "y", "d"},
			array2: []string{"x", "y", "z"},
			output: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := BruteForce(tc.array1, tc.array2)
			if o != tc.output {
				t.Errorf("expected output %t got %v", tc.output, o)
			}
		})
	}
}

func TestUsingHashMap(t *testing.T) {
	tt := []struct {
		name   string
		array1 []string
		array2 []string
		output bool
	}{
		{
			name:   "no match",
			array1: []string{"a", "b", "c", "d"},
			array2: []string{"x", "y", "z"},
			output: false,
		},
		{
			name:   "match",
			array1: []string{"a", "b", "y", "d"},
			array2: []string{"x", "y", "z"},
			output: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := UsingHashMap(tc.array1, tc.array2)
			if o != tc.output {
				t.Errorf("expected output %t got %v", tc.output, o)
			}
		})
	}
}

func BenchmarkBruteForce(b *testing.B) {
	array1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q"}
	array2 := []string{"r", "q", "t", "u", "v", "x", "y", "z"}
	for i := 0; i < b.N; i++ {
		BruteForce(array1, array2)
	}
}

func BenchmarkUsingHashMap(b *testing.B) {
	array1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q"}
	array2 := []string{"r", "q", "t", "u", "v", "x", "y", "z"}
	for i := 0; i < b.N; i++ {
		UsingHashMap(array1, array2)
	}
}
