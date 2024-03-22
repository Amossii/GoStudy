package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	a := person{}
	a.Name = "joe"
	a.Age = 19
	fmt.Println(a)

	b := person{
		Name: "ross",
		Age:  20,
	}
	fmt.Println(b)

	//point
	c := &person{
		Name: "chandler",
		Age:  19,
	}
	c.Age = 100
	fmt.Println(c)
}
