package main

import "fmt"

type (
	byte int8
	str  string
)

func main() {
	var a bool
	var b int64
	var c rune //unicode 字符
	var d float32
	var e float64
	var f [10]int
	var g byte
	var s str = "123"
	fmt.Println(a, b, c, d, e, f, g, s)

}
