package main

import "fmt"

type human struct {
	Name string
}

type person struct {
	human
	Name    string
	Age     int
	contact struct {
		Name string
		Age  int
	}
}

func main() {

	a := struct {
		Name string
		Age  int
	}{
		Name: "Monica",
		Age:  27,
	}
	fmt.Println(a)

	peter := person{
		Name: "peter",
		Age:  19,
		contact: struct {
			Name string
			Age  int
		}{
			Name: "Monica",
			Age:  27,
		}}
	peter.human.Name = "hahaha"
	fmt.Println(peter)
}
