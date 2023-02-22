package gettingstarted

// input -> hash -> output hash number
// insert the data at that output hash number

// what about hash collision
// when two inputs generate the same hashes
// open addressing
// - if taken then insert into the next index
// - will start to lose hash benefits
// separate chaining
// - hash points to a linked list
// - the linked list then contains all the inputs with the
//   same hashed output

// hash takes a key, hashes and returns an int
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)  // convert character to ascii
	}
	return sum % ArraySize
}

// ArraySize is the size of the hash table
const ArraySize = 7

// HashTable will hold the array
type HashTable struct {
	array [ArraySize]*bucket
}

// Insert will take a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if the key is in the hash table.
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take a key and delete it from the hash table.
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket is a linked list
type bucket struct {
	head *bucketNode
}

// insert will take in a key, create a node with the key and insert the
// node in the bucket.
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		//fmt.Println("already exists")
	}
}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}


type bucketNode struct {
	key string
	next *bucketNode
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

