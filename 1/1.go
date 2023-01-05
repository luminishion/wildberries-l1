/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action
	от родительской структуры Human (аналог наследования)
*/

package main

import (
	"fmt"
)

type Human struct {
}

func (h *Human) Privet() {
	fmt.Println("hello")
}

type Action struct {
	Human // embedding
}

func (a *Action) Da() {
	fmt.Println("yes")
}

func main() {
	a := &Action{}
	a.Privet() // call embedded method
	a.Da()
}
