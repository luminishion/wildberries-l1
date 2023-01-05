/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func job(wg *sync.WaitGroup, sum *uint64, val uint64) {
	defer wg.Done()

	atomic.AddUint64(sum, val*val) // add as a single instruction
}

func main() {
	input := [...]uint64{2, 4, 6, 8, 10}
	var sum uint64 = 0

	var wg sync.WaitGroup

	for i := range input {
		wg.Add(1)
		go job(&wg, &sum, input[i])
	}

	wg.Wait()

	fmt.Println(sum)
}
