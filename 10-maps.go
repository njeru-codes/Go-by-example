package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps are like dictioniries in python
	var s = make(map[string]int)

	s["k1"] = 2
	s["k2"] = 34

	fmt.Println("map", s)

	// check if a key exists with error checking
	_, error := s["k2"]
	fmt.Println("error:", error)

	// delete a key pair
	delete(s, "k2")
	fmt.Println("deleted map", s)

	// clear map(delete everything)
	clear(s)
	fmt.Println("empty map", s)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	// check if two maps are equal
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
