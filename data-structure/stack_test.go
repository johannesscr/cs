package data_structure

import "testing"

func TestNewStackNode(t *testing.T) {
	xtc := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "node with string",
			value: "stack node",
		},
		{
			name:  "node with int",
			value: 25,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			node := NewStackNode(tc.value)
			if node.next != nil {
				t.Errorf("expected new stack node to have next pointer to be nil got %v", node.next)
			}
			if node.value != tc.value {
				t.Errorf("expected the stack node to have value %v got %v", tc.value, node.value)
			}
		})
	}
}

func TestNewStack(t *testing.T) {
	xtc := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "no value passed",
			value: nil,
		},
		{
			name:  "initial value",
			value: 5,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStack(tc.value)
			if s.length != 1 {
				t.Errorf("expected stack to have length 1 got %d", s.length)
			}
			if s.top.value != tc.value {
				t.Errorf("expected the stack top to have value %v got %v", tc.value, s.top.value)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	xtc := []struct {
		name        string
		values      []int
		checkFirst  bool
		checkSecond bool
		checkThird  bool
	}{
		{
			name:        "integers",
			values:      []int{5, 7, 2, 9, 8},
			checkFirst:  true,
			checkSecond: true,
			checkThird:  true,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStack(tc.values[0])
			// skip the first one already added at initialization
			for _, v := range tc.values[1:] {
				s.Push(v)
			}

			// test the stack was created correctly
			if s.length != len(tc.values) {
				t.Errorf("expected the stack to have length %d got %d", len(tc.values), s.length)
			}

			if tc.checkFirst {
				val := tc.values[len(tc.values)-1] // last item added
				nodeVal := s.top.value             // last
				if nodeVal != val {
					t.Errorf("expected the first value to be %v got %v", val, nodeVal)
				}
			}
			if tc.checkSecond {
				val := tc.values[len(tc.values)-2] // second last item added
				node := s.top.next                 // second last
				nodeVal := node.value
				if nodeVal != val {
					t.Errorf("expected the first value to be %v got %v", val, nodeVal)
				}
			}
			if tc.checkThird {
				val := tc.values[len(tc.values)-3] // third last item added
				node := s.top.next                 // second last
				node = node.next                   // third last
				nodeVal := node.value
				if nodeVal != val {
					t.Errorf("expected the first value to be %v got %v", val, nodeVal)
				}
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	s := NewStack(5)
	s.Push(8)
	s.Push(3)
	s.Push(5)

	v := s.Pop()
	if v != 5 {
		t.Errorf("expected popped value %v got %v", 5, v)
	}
	v = s.Pop()
	if v != 3 {
		t.Errorf("expected popped value %v got %v", 3, v)
	}
	v = s.Pop()
	if v != 8 {
		t.Errorf("expected popped value %v got %v", 8, v)
	}
	v = s.Pop()
	if v != 5 {
		t.Errorf("expected popped value %v got %v", 5, v)
	}
	// no elements left
	v = s.Pop()
	if v != nil {
		t.Errorf("expected popped value %v got %v", nil, v)
	}
}

func TestStack_Peek(t *testing.T) {
	s := NewStack(8)

	v := s.Peek()
	if v != 8 {
		t.Errorf("expected stack peek to return 8 got %d", v)
	}

	s.Push(19)
	v = s.Peek()
	if v != 19 {
		t.Errorf("expected stack peek to return 19 got %d", v)
	}

	s.Push(12)
	s.Push(6)
	s.Pop()
	v = s.Peek()
	if v != 12 {
		t.Errorf("expected stack peek to return 12 got %d", v)
	}
}

func TestStack_String(t *testing.T) {
	s := NewStack("gophers")
	str := "in --+  +-> out\n     gophers\n"
	if str != s.String() {
		t.Errorf("expected stack string '\n%s\n' got '\n%s\n'", str, s.String())
	}
	s.Push("are")
	str = "in --+  +-> out\n     are\n     gophers\n"
	if str != s.String() {
		t.Errorf("expected stack string '%s' got '\n%s\n'", str, s.String())
	}
	s.Push("here")
	str = "in --+  +-> out\n     here\n     are\n     gophers\n"
	if str != s.String() {
		t.Errorf("expected stack string '%s' got '%s'", str, s.String())
	}
}
