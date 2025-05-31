package main

import (
	"fmt"
)

// errors in Go are communicated via separate return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Test 1: valid divide
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("divide 10, 3:", result)
	}

	// Test 2: divide by zero
	newResult, newErr := divide(10, 0)
	if newErr != nil {
		fmt.Println("new error:", newErr)
	} else {
		fmt.Println("new result:", newResult)
	}
}
