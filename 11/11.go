/*
Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import (
	"fmt"
)

func main() {
	v1 := []int{5, 2, 8, 7, 4}
	v2 := []int{1, 9, 6, 4, 7}

	m := make(map[int]bool)
	for _, v := range v1 {
		m[v] = true // mark in map that first array has such value
	}

	var ret []int
	for _, v := range v2 {
		if m[v] { // add to ret if marked
			ret = append(ret, v)
		}
	}

	fmt.Println(ret)
}
