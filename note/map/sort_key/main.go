package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "C", 4: "d", 5: "e"}
	s := make([]int, len(m))

	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}

	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
}
