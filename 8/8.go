/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

package main

import (
	"fmt"
)

func setIBit(num int64, i int, val bool) int64 {
	var flag int64 = 1 << (i - 1) // all 0, i-bit is 1

	if val {
		return num | flag // add flag
	} else {
		return num &^ flag // remove flag
	}
}

func swapIBit(num int64, i int) int64 {
	var flag int64 = 1 << (i - 1)
	return num ^ flag // change flag
}

func main() {
	var num int64 = 255
	i := 3

	fmt.Printf("num =        	%b\n", num)

	num = setIBit(num, i, false)
	fmt.Printf("make bit %d = 0	%b\n", i, num)

	num = setIBit(num, i, true)
	fmt.Printf("make bit %d = 1	%b\n", i, num)

	num = swapIBit(num, i)
	fmt.Printf("make bit %d xor	%b\n", i, num)

	num = swapIBit(num, i)
	fmt.Printf("make bit %d xor	%b\n", i, num)
}
