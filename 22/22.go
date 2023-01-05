/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Float)
	a.SetString("99999999999999999999")

	b := new(big.Float)
	b.SetString("99999999999999999999")

	c := new(big.Float)

	c.Quo(a, b)
	fmt.Println(c)

	c.Add(a, b)
	fmt.Println(c)

	c.Mul(a, b)
	fmt.Println(c)

	c.Sub(a, b)
	fmt.Println(c)
}
