package main

import "fmt"

func main() {
	c := C
	c()

	//anonymous funciton
	a := func() {
		fmt.Println("this is an anonymous function")
	}
	a()
}

// declare a function
func A(a int, b int, c int) (e, f, d int) {
	return e, f, d
}

func B(a ...int) {
	//a is a slice
	fmt.Println(a)
}

func C() {
	fmt.Println("this is function C")
}
