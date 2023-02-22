package data_structure

import (
	"errors"
	"fmt"
	"testing"
)

func TestCustomHash(t *testing.T) {
	d := []byte("name")
	CustomHash(d)
	d = []byte("Hello")
	CustomHash(d)
}
func TestBasicHash(t *testing.T) {
	h1 := BasicHash("hello", 5)
	if h1 != 0 {
		t.Errorf("expected hash for 'hello' to be 0 got %d", h1)
	}
	h2 := BasicHash("name", 4)
	if h2 != 2 {
		t.Errorf("expected hash for 'name' to be 2 got %d", h2)
	}
}

func TestNewBasicHashTable(t *testing.T) {
	bht := NewBasicHashTable(5)
	if bht.size != 5 {
		t.Errorf("expected basic hash table with size 5 got %d", bht.size)
	}
	h1 := bht.hash("hello")
	if h1 != 0 {
		t.Errorf("expected hash for 'hello' to be 0 got %d", h1)
	}
	if len(bht.data) != 5 {
		t.Errorf("expected basic hash table to have initial date length of 5 got %d", len(bht.data))
	}
}

func TestBasicHashTable_Set(t *testing.T) {
	bht := NewBasicHashTable(4)
	bht.Set("name", 125)
	if len(bht.data[2]) == 0 {
		t.Errorf("expected bucket not to be empty empty got %v", len(bht.data[2]))
	}
	hp := HashPair{
		Key:   "name",
		Value: 125,
	}
	if bht.data[2][0] != hp {
		t.Errorf("expected to have a hash pair %v got %v", hp, bht.data[2][0])
	}
	bht.Set("james", 132)
	hp = HashPair{
		Key:   "james",
		Value: 132,
	}
	if bht.data[2][1] != hp {
		t.Errorf("expected to have a hash pair %v got %v", hp, bht.data[2][1])
	}
}

func TestBasicHashTable_Get(t *testing.T) {
	bht := NewBasicHashTable(4)
	bht.Set("name", 125)  // 2
	bht.Set("james", 132) // 2
	bht.Set("ken", 871)   // 1
	num, err := bht.Get("ken")
	if num != 871 {
		t.Errorf("expected value %d got %d", 871, num)
	}
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	num, err = bht.Get("james")
	if num != 132 {
		t.Errorf("expected value %d got %d", 132, num)
	}
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	num, err = bht.Get("name")
	if num != 125 {
		t.Errorf("expected value %d got %d", 125, num)
	}
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	num, err = bht.Get("franky")
	if num != 0 {
		t.Errorf("expected value %d got %d", 0, num)
	}
	e := errors.New("no hash pair found")
	if err.Error() != e.Error() {
		t.Errorf("expected err to %v got %v", e, err)
	}
}

func TestBasicHashTable_Keys(t *testing.T) {
	bht := NewBasicHashTable(4)
	bht.Set("grapes", 125)
	bht.Set("apples", 132)
	bht.Set("oranges", 871)
	bht.Set("mangoes", 712)
	xKeys := bht.Keys()
	xKeysString := "[oranges mangoes grapes apples]"
	if len(xKeys) != 4 {
		t.Errorf("expected keys to have length 4 got %d", len(xKeys))
	}
	if xKeysString != fmt.Sprintf("%v", xKeys) {
		t.Errorf("expected keys '%s' got '%v'", xKeysString, xKeys)
	}
}
