package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type manager struct {
	User
	title string
}

func main() {
	m := manager{User: User{1, "ok", 2}, title: "manager"}
	t := reflect.TypeOf(m)
	fmt.Printf("%v\n", t.Field(1))
	//取出User的字段
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 1}))
}
