package main

import "fmt"

// try to change a map's key to value,and value to key
func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m)

	rm := make(map[string]int)
	for k, v := range m {
		rm[v] = k
	}
	fmt.Println(rm)
}
