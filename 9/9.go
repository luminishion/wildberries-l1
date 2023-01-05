/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные
из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"os"
)

func twicer(input, output chan int) {
	for val := range input {
		output <- val * 2
	}

	close(output)
}

func writer(input chan int, m []int) {
	for _, val := range m {
		input <- val
	}

	close(input)
}

func reader(output chan int) {
	for val := range output {
		fmt.Fprintln(os.Stdout, val)
	}
}

func main() {
	buf := 3

	input := make(chan int, buf)
	output := make(chan int, buf)
	go twicer(input, output)

	m := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go writer(input, m[:])

	reader(output)
}
