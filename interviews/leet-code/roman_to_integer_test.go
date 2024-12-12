package leet_code

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "III",
			out: 3,
		},
		{
			in:  "LVIII",
			out: 58,
		},
		{
			in:  "MCMXCIV",
			out: 1994,
		},
	}

	for _, tc := range tests {
		t.Run("roman-to-integer", func(t *testing.T) {
			o := romanToInt(tc.in)
			if o != tc.out {
				t.Errorf("expected output %s: %d got %d", tc.in, tc.out, o)
			}
		})
	}
}
