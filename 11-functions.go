package main

import "fmt"

// return singel value
func addValues(a int, b int) int {
	return a + b
}

// return mutiple values
func vals() (int, int) {
	return 3, 6
}

// Variadic Functions --> can be called with any number of args
func summation(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum := addValues(1, 2) //call the function
	fmt.Println("summation", sum)

	// mutiple return values
	a, b := vals()
	fmt.Println(a, b)

	// ignore one value
	_, c := vals()
	fmt.Println(c)

	// Variadic Functions
	summation(1, 2)
	summation(1, 200, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	summation(nums...)

}
