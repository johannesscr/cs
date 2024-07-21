package main

import "testing"

func TestMemo(t *testing.T) {
	Memo(5)
	Memo(5)
	Memo(5)
}

func TestMemo2(t *testing.T) {
	memoized := Memo2()
	memoized(5)
	memoized(5)
	memoized(5)
}

func TestNoMemo(t *testing.T) {
	NoMemo(5)
	NoMemo(5)
	NoMemo(5)
}

func TestFib(t *testing.T) {
	Fib(15)
}

func TestFibNormal(t *testing.T) {
	WrapFibNormal(15)
}
