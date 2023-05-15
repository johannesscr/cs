package challenges

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{
			in:  121,
			out: true,
		},
		{
			in:  -121,
			out: false,
		},
		{
			in:  12321,
			out: true,
		},
		{
			in:  1221,
			out: true,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			o := IsPalindrome(tc.in)
			if o != tc.out {
				t.Errorf("expected %t for %d got %t", tc.out, tc.in, o)
			}
		})
	}
}