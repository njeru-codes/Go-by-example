package main

import "fmt"

func main() {
	var i = 1

	// normal loops
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// c- like loops
	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}

	// python-like loop
	for k := range 5 {
		fmt.Println("range", k)
	}

	// break
	for {
		fmt.Println("loop")
		break
	}

	// continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
