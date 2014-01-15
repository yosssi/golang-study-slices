package main

import (
	"fmt"
	"io/ioutil"
)

const (
	filePath string = "./sample.txt"
)

func main() {

	// Read the file and get a byte slice.
	b, _ := ioutil.ReadFile(filePath)

	// Re-slice the byte slice.
	s := b[:10]

	// Length of the slice is short but capacity of the slice is long.
	print("Before", s)

	// Create a new slice.
	s = append([]byte{}, s...)

	// Both length and capacity of the slice are short.
	print("After ", s)
}

func print(label string, s []byte) {
	fmt.Printf("[%s] content: %s, len: %d, cap: %d\n", label, s, len(s), cap(s))
}
