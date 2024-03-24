package main

import "fmt"

type A struct {
	name string //同一个包都可见
	Age  int    //别的包都可见
}

func main() {
	a := A{name: "peter", Age: 18}
	fmt.Println(a.name)
	fmt.Println(a.Age)
	a.ChangeName()
	fmt.Println(a.name)
}
func (a *A) ChangeName() {
	a.name = "monica"

}
