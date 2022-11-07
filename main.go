package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
)

// ReadData takes a file name for a specific data structure. The data
// structures are stored as JSON objects, with the structs included in
// the file called types.go from the package main.
//
// see the list of filenames
func ReadData(filename string) []byte {
	// get the working directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get working directory")
	}
	// create the path from the working directory and the filename
	p := path.Join(pwd, filename)
	// read the file into a slice of bytes
	xb, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalf("Unable to read file with path: %s", p)
	}
	// return the data as a slice of bytes
	return xb
}

// UnmarshalData simply does the error handling for when the data needs to
// be unmarshalled.
func UnmarshalData(xb []byte, v interface{}) {
	err := json.Unmarshal(xb, v)
	if err != nil {
		log.Fatal("Unable to unmarshal data")
	}
}

// MarshallAndWrite marshall's the data to JSON format, then writes the
// marshalled data to a file based on the filename.
func MarshallAndWrite(xb []byte, filename string) bool {
	switch filename {
	case "int-array.json":
		//xb, err := json./Marshal()
	}
	return true
}

func main() {
	x := make([]int, 10)
	for i := 0; i < 1000000; i++ {
		r := rand.Intn(1000)
		x = append(x, r)
	}
	fmt.Println(x)
}
