/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
)

func main() {
	var v interface{} = make(chan int)
	t := fmt.Sprintf("%T", v)
	fmt.Println(t)
}
