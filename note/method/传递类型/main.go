package main

import "fmt"

//函数传递为值传递,如果需要传递引用需要加*，而且作为结构体无需加&
type A struct {
	Name string
}

func main() {
	a := A{Name: "this is a "}
	fmt.Println(a.Name)
	a.print()
	fmt.Println(a.Name)
	a.print2()
	fmt.Println(a.Name)
}
func (a A) print() {
	a.Name = "this is a new A"
	fmt.Println(a.Name)
}
func (a *A) print2() {
	a.Name = "this is a new A"
	fmt.Println(a.Name)
}
