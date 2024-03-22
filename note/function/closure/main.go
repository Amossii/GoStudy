package main

import "fmt"

func main() {
	f := Closure(10)
	fmt.Println(f(200))
	fmt.Println(f(100))
}
func Closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
