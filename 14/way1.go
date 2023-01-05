/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = make(chan int)
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("chan")
	default:
		fmt.Println("no case")
	}
}
