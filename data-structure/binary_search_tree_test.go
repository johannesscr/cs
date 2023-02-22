package data_structure

import (
	"encoding/json"
	"os"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	bs := NewBinarySearchTree(5)
	if bs.Root.Value != 5 {
		t.Errorf("expected binary search tree with root value 5 got %d", bs.Root.Value)
	}
	if bs.Root.Left != nil {
		t.Errorf("expected binary search tree root left to be nil got %v", bs.Root.Left)
	}
	if bs.Root.Right != nil {
		t.Errorf("expected binary search tree root left to be nil got %v", bs.Root.Right)
	}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := NewBinarySearchTree(9)
	bst.Insert(4)
	bst.Insert(20)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(170)
	bst.Insert(1)
	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(bst)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(bst.Root, bst.root.left, bst.root.right)
	//left := bst.Root.left
	//fmt.Println(left, left.left, left.right)
}

func TestBinarySearchTree_Lookup(t *testing.T) {
	bst := NewBinarySearchTree(9)
	bst.Insert(4)
	bst.Insert(20)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(170)
	bst.Insert(1)

	n := bst.Lookup(4)
	if n == nil {
		t.Errorf("expected node not nil got %v", n)
	}

	n = bst.Lookup(6)
	if n == nil {
		t.Errorf("expected node not nil got %v", n)
	}

	n = bst.Lookup(16)
	if n != nil {
		t.Errorf("expected node to be nil got %v", n)
	}
}
