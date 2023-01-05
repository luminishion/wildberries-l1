/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

package main

import (
	"fmt"
)

func reverse(str string) string {
	buf := []rune(str)

	// i, j = j, i until the middle
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}

func main() {
	fmt.Println(reverse("абвг6"))
}
