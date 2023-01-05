/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	v int64
}

func (c *Counter) Add() int64 {
	return atomic.AddInt64(&c.v, 1)
}

func (c *Counter) Get() int64 {
	return atomic.LoadInt64(&c.v)
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
