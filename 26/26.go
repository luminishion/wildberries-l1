/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	s = strings.ToLower(s)

	m := make(map[rune]bool)

	for _, v := range []rune(s) {
		if m[v] { // if already marked
			return false
		}

		m[v] = true // mark
	}

	return true
}

func main() {
	fmt.Println(check("abcd"))
	fmt.Println(check("abCdefAaf"))
	fmt.Println(check("aabcd"))
}
