package main

import "fmt"

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GOGOGO!")
		c <- true
	}()
	<-c
}
func Go() {
	fmt.Println("GOGOGO!")
}
