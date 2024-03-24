package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (user User) Hello() string {
	fmt.Println("Hello world")
	return "Hello world"
}
func main() {
	u := User{1, "ok", 12}
	info(u)
}
func info(o interface{}) {
	//判断类型
	t := reflect.TypeOf(o)
	fmt.Println("type:", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("field:")
	//获取属性
	//     ID:int=1
	//   Name:string=ok
	//    Age:int=12
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n", f.Name, f.Type, val)
	}
	//打印方法
	// Hello:func(main.User) string
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}

}
