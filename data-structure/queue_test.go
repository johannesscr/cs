package data_structure

import "testing"

func TestNewQueueNode(t *testing.T) {
	xtc := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "string queue",
			value: "gopher",
		},
		{
			name:  "integer queue",
			value: 21,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			n := NewQueueNode(tc.value)
			if n.value != tc.value {
				t.Errorf("expected new queue node to be %v got %v", tc.value, n.value)
			}
			if n.next != nil {
				t.Errorf("expected a new queue node to point to a next node of nil got %v", n.next)
			}
		})
	}
}

func TestNewQueue(t *testing.T) {
	xtc := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "string queue",
			value: "original gopher",
		},
		{
			name:  "integer queue",
			value: 12,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue(tc.value)
			if q.length != 1 {
				t.Errorf("expected initial queue to have length 1 got %d", q.length)
			}
			if q.start.value != tc.value {
				t.Errorf("expected initial queue start node to have value %v got %v", q.start.value, tc.value)
			}
			if q.end.value != tc.value {
				t.Errorf("expected initial queue end node to have value %v got %v", q.end.value, tc.value)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	xtc := []struct {
		name             string
		values           []int
		checkStartFirst  bool
		checkStartSecond bool
		checkStartThird  bool
	}{
		{
			name:             "queue of ints",
			values:           []int{5, 7, 13, 9, 4},
			checkStartFirst:  true,
			checkStartSecond: true,
			checkStartThird:  true,
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue(tc.values[0])
			// add all the queue entries
			for _, v := range tc.values[1:] {
				q.Enqueue(v)
			}

			if tc.checkStartFirst {
				v := q.start.value
				if v != tc.values[0] {
					t.Errorf("expected queue start first value to be %v got %v", tc.values[0], v)
				}
			}
			if tc.checkStartSecond {
				node := q.start.next // second in queue
				v := node.value
				if v != tc.values[1] {
					t.Errorf("expected queue start first value to be %v got %v", tc.values[1], v)
				}
			}
			if tc.checkStartThird {
				node := q.start.next // second in queue
				node = node.next     // third in queue
				v := node.value
				if v != tc.values[2] {
					t.Errorf("expected queue start first value to be %v got %v", tc.values[2], v)
				}
			}
			endOfQueueValue := q.end.value
			endValue := tc.values[len(tc.values)-1]
			if endOfQueueValue != endValue {
				t.Errorf("expected the end of the queue to have value %v got %v", endValue, endOfQueueValue)
			}
		})
	}
}

func TestQueue_String(t *testing.T) {
	xtc := []struct {
		name   string
		values []int
		output string
	}{
		{
			name:   "new queue",
			values: []int{5},
			output: "in -> 5 -> out",
		},
		{
			name:   "small queue",
			values: []int{5, 12, 7, 4},
			output: "in -> 5 -> 12 -> 7 -> 4 -> out",
		},
		{
			name:   "longer queue",
			values: []int{5, 12, 7, 4, 9, 5, 7},
			output: "in -> 5 -> 12 -> 7 -> 4 -> 9 -> 5 -> 7 -> out",
		},
	}

	for _, tc := range xtc {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue(tc.values[0])
			for _, v := range tc.values[1:] {
				q.Enqueue(v)
			}

			output := q.String()
			if output != tc.output {
				t.Errorf("expected a queue representation string '%s' got '%s'", tc.output, output)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue(8)
	q.Enqueue(3)
	q.Enqueue(5)

	v := q.Dequeue()
	if v != 8 {
		t.Errorf("expected dequeue to remove 8 got %v", v)
	}
	if q.length != 2 {
		t.Errorf("expected queue to have length 2 got %d", q.length)
	}
	v = q.Dequeue()
	if v != 3 {
		t.Errorf("expected dequeue to remove 3 got %v", v)
	}
	if q.length != 1 {
		t.Errorf("expected queue to have length 1 got %d", q.length)
	}
	v = q.Dequeue()
	if v != 5 {
		t.Errorf("expected dequeue to remove 5 got %v", v)
	}
	if q.length != 0 {
		t.Errorf("expected queue to have length 0 got %d", q.length)
	}
	v = q.Dequeue()
	if v != nil {
		t.Errorf("expected dequeue to remove nil got %v", v)
	}
	if q.length != 0 {
		t.Errorf("expected queue to have length 0 got %d", q.length)
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue(5)
	if q.Peek() != 5 {
		t.Errorf("expected queue peek to be 5 got %v", q.Peek())
	}
	if q.length != 1 {
		t.Errorf("expected queue to have length of 1 got %v", q.length)
	}

	q.Enqueue(8)
	if q.Peek() != 5 {
		t.Errorf("expected queue peek to be 5 got %v", q.Peek())
	}
	if q.length != 2 {
		t.Errorf("expected queue to have length of 2 got %v", q.length)
	}

	q.Dequeue()
	if q.Peek() != 8 {
		t.Errorf("expected queue peek to be 8 got %v", q.Peek())
	}
	if q.length != 1 {
		t.Errorf("expected queue to have length of 1 got %v", q.length)
	}
}
