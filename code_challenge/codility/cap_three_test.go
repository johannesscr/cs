package codility

import "testing"

func capThree() bool {
	return false
}

func TestCapThree(t *testing.T) {
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
			out := capThree()
			if test.output != out {
				t.Errorf("Expected %v, got %v", test.output, out)
			}
		})
	}
}
