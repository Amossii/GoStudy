package main

import "fmt"

var (
	a int
	b bool
)

func main() {
	c := 10 //全局变量不可使用
	a = 10
	b = true
	fmt.Println(a, b, c)
}
