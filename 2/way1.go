/*
	Написать программу, которая конкурентно рассчитает
	значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"os"
	"sync"
)

func job(wg *sync.WaitGroup, val *float64) {
	defer wg.Done() // decr wait counter

	*val = *val * *val
}

func main() {
	input := [...]float64{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := range input {
		wg.Add(1)              // incr wait counter
		go job(&wg, &input[i]) // run job async
	}

	wg.Wait() // wait until all jobs done

	for _, v := range input {
		fmt.Fprintln(os.Stdout, v)
	}
}
