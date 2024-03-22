package main

import "fmt"

func main() {
	// initial
	var m1 map[int]string = map[int]string{}
	fmt.Println(m1)

	m2 := make(map[int]string)
	fmt.Println(m2)

	//key
	m2[0] = "ok"
	m2[1] = "hello"
	m2[2] = "world"
	a := m2[2]
	fmt.Println(a)

	//delete
	fmt.Println("before delete:", m2)
	delete(m2, 0)
	fmt.Println("after delete:", m2)

	//value type is map
	var m map[int]map[int]string = make(map[int]map[int]string)
	m[1] = m2
	fmt.Println(m)

	//map in slice
	sm := make([]map[int]string, 5)
	for i, _ := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][i] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
