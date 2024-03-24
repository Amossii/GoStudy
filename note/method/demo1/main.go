package main

import "fmt"

type A struct {
	Name string
}
type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()

	b := B{}
	b.Print()
}
func (a A) Print() {
	fmt.Println("this is struct A")
}
func (a B) Print() {
	fmt.Println("this is struct B")
}
