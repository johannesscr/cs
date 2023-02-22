package data_structure

import (
	"crypto/md5"
	"errors"
	"fmt"
)

func CustomHash(data []byte) {
	fmt.Printf("hash '%v' => MD5 => %x\n", data, md5.Sum(data))
}

func BasicHash(key string, length int) int {
	hash := 0
	for i, s := range key {
		hash = (hash + int(s)*i) % length
	}
	return hash
}

type HashPair struct {
	Key   string
	Value int
}

type BasicHashTable struct {
	data [][]HashPair
	size int
}

func NewBasicHashTable(size int) BasicHashTable {
	h := BasicHashTable{
		size: size,
		data: make([][]HashPair, size),
	}
	return h
}

// hash is the function used by the BasicHashTable to create the hash address
// for the bucket in which key value pair is to be added. It takes the key and
// returns the integer address.
//
// We consider this to be really fast: O(1)
func (bht *BasicHashTable) hash(key string) int {
	return BasicHash(key, bht.size)
}

// Set adds the new key-value pair to the BasicHashTable.
//
// We simply push/append on the end of the slice: O(1)
func (bht *BasicHashTable) Set(key string, value int) {
	// would be good to check the duplicate keys cannot be added.

	// create the hash address
	address := bht.hash(key)
	// create the hash pair
	hp := HashPair{
		Key:   key,
		Value: value,
	}
	// add the hash pair in the allocated address
	bht.data[address] = append(bht.data[address], hp)
}

// Get retrieves the value for the given key in the BasicHashTable as well as an
// error if there is key-value pair cannot be found.
//
// Most of the time, with no collisions it is O(1), but is a bad example this
// can be O(n) worst case. However, with size large enough, this should not
// occur.
func (bht *BasicHashTable) Get(key string) (int, error) {
	// create the hash address
	address := bht.hash(key)

	// since the bucket may have had collisions, we need to search through all
	// possible collision values to find the correct key-value pair
	for _, hp := range bht.data[address] {
		if hp.Key == key {
			return hp.Value, nil
		}
	}
	return 0, errors.New("no hash pair found")
}

// Keys returns all the keys in the BasicHashTable.
//
// Worst case we are looking at an O(n^2)
func (bht *BasicHashTable) Keys() []string {
	xKeys := make([]string, 0)
	for _, buckets := range bht.data { // O(size) the hash table size
		for _, hp := range buckets { // O(n/size)
			xKeys = append(xKeys, hp.Key)
		}
	}
	return xKeys
}
