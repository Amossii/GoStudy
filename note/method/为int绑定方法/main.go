package main

import "fmt"

type TZ int

func main() {
	var a TZ
	a.print()
	(*TZ).print(&a)
}
func (t *TZ) print() {
	fmt.Println(t)
	fmt.Println("this is TZ's method!")
}
