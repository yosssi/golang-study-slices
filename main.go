package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

const (
	filePath string = "./sample.txt"
)

func main() {
	digitRegexp := regexp.MustCompile("[0-9]+")

	// Read the file and get a byte slice.
	b, _ := ioutil.ReadFile(filePath)

	// Re-slice the byte slice by getting the consecutive numeric digits.
	s := digitRegexp.Find(b)

	// Length of the slice is short but Capacity of the slice is long.
	print("Before", s)

	// Create a new slice.
	s = append([]byte{}, s...)

	// Both length and capaticy of the slice are short.
	print("After ", s)
}

func print(label string, s []byte) {
	fmt.Printf("[%s] content: %s, len: %d, cap: %d\n", label, s, len(s), cap(s))
}
