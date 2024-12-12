package gettingstarted

import (
	"testing"
)

func TestBucket(t *testing.T) {
	tb := &bucket{}
	tb.insert("randy")
	tb.insert("randy")
	//fmt.Println(tb.search("randy"))
	tb.delete("randy")
	//fmt.Println(tb.search("randy"))
	//fmt.Println(tb.search("eric"))
}
