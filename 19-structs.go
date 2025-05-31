package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 34

	return &p
}
func main() {
	p := newPerson("alice")
	fmt.Println(p)
}
