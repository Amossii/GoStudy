package main

import "fmt"

type A struct {
	Name string
}

func main() {
	a := A{Name: "My name is A"}
	// a.print()
	fmt.Println(&a)
	(*A).print(&a)
	fmt.Println(a.Name)

	(A).print2(a)
}
func (a *A) print() {
	fmt.Println(a)
	a.Name = "this is B"
	fmt.Println(a.Name)
}
func (a A) print2() {
	fmt.Println(&a)
	fmt.Println(a.Name)
}
