package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

type Marker string

func (m Marker) Valid() bool {
	for i1, r1 := range m {
		for i2, r2 := range m {
			if i1 != i2 && r1 == r2 {
				return false
			}
		}
	}

	return true
}

func findMarker(length int) int {
	for i := length; i < len(input); i++ {
		if marker := Marker(input[i-length : i]); marker.Valid() {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Printf("found 4er marker: %d\n", findMarker(4))
	fmt.Printf("found 14er marker: %d\n", findMarker(14))
}
