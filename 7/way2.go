/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.RWMutex

	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			m[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(m)
}
