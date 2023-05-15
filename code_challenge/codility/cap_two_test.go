package codility

import "testing"

func capTwo() bool {
	return false
}

func TestCapTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  interface{}
		output interface{}
	}{
		{
			name: "test one",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := capTwo()
			if test.output != out {
				t.Errorf("Expected %v, got %v", test.output, out)
			}
		})
	}
}
