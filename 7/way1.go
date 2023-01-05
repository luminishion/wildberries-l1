/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{} // thread safe map
	var wg sync.WaitGroup

	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			m.Store(i, i)
		}()
	}

	wg.Wait()
	fmt.Println(m)
}
