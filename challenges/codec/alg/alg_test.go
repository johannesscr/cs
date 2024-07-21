package alg

import "testing"

func TestRobotDirectionAfter(t *testing.T) {
	testcases := []struct {
		name     string
		distance float32
		output   string
	}{
		{
			name:     "init",
			distance: 0,
			output:   "N",
		},
		{
			name:     "partial",
			distance: 0.05,
			output:   "N",
		},
		{
			name:     "one",
			distance: 0.3,
			output:   "S",
		},
		{
			name:     "two",
			distance: 7.25,
			output:   "N",
		},
		{
			name:     "three",
			distance: 1.2,
			output:   "E",
		},
		{
			name:     "four",
			distance: 19.62,
			output:   "W",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			finalDirection := robotDirectionAfter(tc.distance)
			if finalDirection != tc.output {
				t.Errorf("expected direction %s got %s", tc.output, finalDirection)
			}
		})
	}
}

func TestMineSweeper(t *testing.T) {
	m := []string{"XOO", "OOO", "XXO"}
	out := "X 1 0\n3 3 1\nX X 1"
	o := minesweeper(m)
	if o != out {
		t.Errorf("expected \n%s\n got \n%s\n", out, o)
	}
}
