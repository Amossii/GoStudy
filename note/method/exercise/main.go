package main

import "fmt"

type TZ int

func main() {
	var a TZ = 0
	fmt.Println(a)
	a.Increase()
	fmt.Println(a)
}
func (t *TZ) Increase() {
	*t += 100
}
