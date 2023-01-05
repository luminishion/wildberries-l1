/*
	Написать программу, которая конкурентно рассчитает
	значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"os"
)

func job(val float64, output chan float64) {
	output <- val * val // write result to channel
	// just print here would be unfair
}

func main() {
	input := [...]float64{2, 4, 6, 8, 10}
	output := make(chan float64, len(input))
	// buffer size have to be same as goroutines count otherwise some of them may have to wait until channel is free

	for _, v := range input {
		go job(v, output)
	}

	for range input {
		fmt.Fprintln(os.Stdout, <-output) // read from channel, gonna freeze until data arrives
	}
}
