package main

import (
	"fmt"
	"io/ioutil"
)

const (
	filePath      string = "./sample.txt"
	sliceEndPoint int    = 10
)

func main() {

	// Re-slice the byte slice.
	s := reslice()

	// Length of the slice is short but capacity of the slice is long.
	print("Before", s)

	// Create a new slice from the byte slice.
	s = create()

	// Both length and capacity of the slice are short.
	print("After ", s)
}

func reslice() []byte {
	// Read the file and get a byte slice.
	b, _ := ioutil.ReadFile(filePath)

	// Re-slice the byte slice.
	s := b[:sliceEndPoint]

	// Return the slice.
	return s
}

func create() []byte {
	// Re-slice the byte slice.
	s := reslice()

	// Create a new slice from the byte slice.
	s = append([]byte{}, s...)

	// Return the slice.
	return s
}

func print(label string, s []byte) {
	fmt.Printf("[%s] content: %s, len: %d, cap: %d\n", label, s, len(s), cap(s))
}
