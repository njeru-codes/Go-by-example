package main

import "fmt"

// pinter  allow you to pass references to
// values and records within your program

func changeValue(x *int) {
	*x = 100
}

func main() {

	a := 10
	changeValue(&a)
	fmt.Println(a)

}

// when to use pointers in go
/**
1. Passing a large struct to a function can be expensive. Use a pointer to avoid copying.
2. Go passes arguments by value, so if you want to modify the original variable, use a pointer
3. Pointers are essential for building structures where elements reference each other.
4. Go doesn’t have null variables for primitives, but you can use pointers to indicate absence.
*/

// real life use case
/*
1. Config structs	Modify shared config in memory
2. JSON unmarshaling
3. Networking (TCP/UDP structs)	Manage connection state via pointers
4. Resource management	Manage pooled resources with shared ownership
*/

/**
Summary: When to Use Pointers

Use pointers in Go when:

    You need to modify a value inside a function.

    You want to avoid copying large data.

    You want to indicate a missing or optional value.

    You’re working with complex data structures.

*/
