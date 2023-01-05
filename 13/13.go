/*
Поменять местами два числа без создания временной переменной.
*/

package main

import (
	"fmt"
)

func main() {
	a, b := 1, 0
	b, a = a, b
	fmt.Println(a, b)
}
