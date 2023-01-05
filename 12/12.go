/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

package main

import (
	"fmt"
)

func main() {
	v := []string{"cat", "cat", "dg", "cat", "tree"}

	m := make(map[string]bool)
	for _, v := range v {
		m[v] = true // mark all values of v
	}

	var ret []string
	for v := range m {
		ret = append(ret, v) // all keys to array
	}

	fmt.Println(ret)
}
