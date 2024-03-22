// this is a example of panic
package main

import "fmt"

func main() {
	A()
	B()
	C()
}
func A() {
	fmt.Println("this is function A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}
func C() {
	fmt.Println("this is function C")
}
