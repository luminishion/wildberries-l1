/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	p := &Point{
		x: x,
		y: y,
	}

	return p
}

func (p *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2))
}

func (p *Point) Pos() (float64, float64) {
	return p.x, p.y
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(0, 0)
	fmt.Println(p1.Distance(p2))
}
