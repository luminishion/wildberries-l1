/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

package main

import (
	"fmt"
)

func main() {
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	ret := make(map[int][]float64)

	for _, v := range input {
		id := int(v/10) * 10         // can't use math.floor because of sign
		ret[id] = append(ret[id], v) // group by same id
	}

	fmt.Println(ret)
}
