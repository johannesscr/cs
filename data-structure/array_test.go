package data_structure

import "testing"

func TestNewArray(t *testing.T) {
	a := NewArray()
	if len(a.data) != 0 {
		t.Errorf("expected new array to have no data got %v", a.data)
	}
	if a.length != 0 {
		t.Errorf("expected new array to have length 0 got %d", a.length)
	}
}

func TestArray_Push(t *testing.T) {
	a := NewArray()
	if len(a.data) != 0 {
		t.Errorf("expected new array to have no data got %v", a.data)
	}
	a.Push(1)
	if len(a.data) == 0 {
		t.Errorf("expected array data to have an element after push got %v", a.data)
	}
	if a.length != 1 {
		t.Errorf("expected array to have length of 1 got %d", a.length)
	}
}

func TestArray_Get(t *testing.T) {
	a := NewArray()
	v := a.Get(2)
	if v != nil {
		t.Errorf("expected out of index to return nil got %v", v)
	}
	a.Push(1)
	v = a.Get(0)
	if v != 1 {
		t.Errorf("expected pushed element 1 got %v", v)
	}
}

func TestMergeSortedArrays(t *testing.T) {
	a := []int{0, 3, 4, 31}
	b := []int{4, 6, 30}
	output := []int{0, 3, 4, 4, 6, 30, 31}

	o := MergeSortedArrays(a, b)
	for i, v := range output {
		oi := o[i]
		if oi != v {
			t.Errorf("expected array element %d in position %d got %d", v, i, oi)
		}
	}
}

//func TestMergeSortedArrayExample(t *testing.T) {
//	a := []int{0, 3, 4, 31}
//	b := []int{4, 6, 30}
//	output := []int{0, 3, 4, 4, 6, 30, 31}
//
//	o := MergeSortedArrayExample(a, b)
//	for i, v := range output {
//		oi := o[i]
//		if oi != v {
//			t.Errorf("expected array element %d in position %d got %d", v, i, oi)
//		}
//	}
//}
