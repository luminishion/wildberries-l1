/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

func job(wg *sync.WaitGroup, mu *sync.RWMutex, sum *uint64, val uint64) {
	defer wg.Done()

	mu.Lock()         // make sure another goroutines do not change sum at same time
	defer mu.Unlock() // unlock access to this section

	*sum += val * val
}

func main() {
	input := [...]uint64{2, 4, 6, 8, 10}
	var sum uint64 = 0

	var wg sync.WaitGroup
	var mu sync.RWMutex

	for i := range input {
		wg.Add(1)
		go job(&wg, &mu, &sum, input[i])
	}

	wg.Wait()

	fmt.Println(sum)
}
