/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	buf := strings.Fields(str)

	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return strings.Join(buf, " ")
}

func main() {
	fmt.Println(reverse("vam povestka"))
}
