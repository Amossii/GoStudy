package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world ")
	//并发
	go Go()
	time.Sleep(2 * time.Second)

}
func Go() {
	fmt.Println("go go go")
}
