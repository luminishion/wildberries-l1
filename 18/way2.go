/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.RWMutex
	v int64
}

func (c *Counter) Add() {
	c.Lock()
	defer c.Unlock()

	c.v++
}

func (c *Counter) Get() int64 {
	c.RLock()
	defer c.RUnlock()

	return c.v
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Add()
		}()
	}
	wg.Wait()

	fmt.Println(c.Get())
}
