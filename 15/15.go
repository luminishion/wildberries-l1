/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import (
	"strings"
)

func createHugeString(len int) string {
	return strings.Repeat("A", len)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100]) // without copy garbage collector can not collect linked huge string
}

func main() {
	someFunc()
}
