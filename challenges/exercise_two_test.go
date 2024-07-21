package challenges

import (
	"testing"
)

func TestRecurringInt(t *testing.T) {
	tt := []struct {
		name   string
		xi     []int
		output int
		ok     bool
	}{
		{
			name:   "should return 2",
			xi:     []int{2, 5, 1, 2, 3, 5, 1, 2, 4},
			output: 2,
			ok:     true,
		},

		{
			name:   "should return 1",
			xi:     []int{2, 1, 1, 2, 3, 5, 1, 2, 4},
			output: 1,
			ok:     true,
		},

		{
			name:   "should return nil",
			xi:     []int{2, 3, 4, 5},
			output: 0,
			ok:     false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o, ok := RecurringInt(tc.xi)
			if o != tc.output {
				t.Errorf("expected output %d got %d", tc.output, o)
			}
			if ok != tc.ok {
				t.Errorf("expected ok %t got %t", tc.ok, ok)
			}
		})
	}
}
