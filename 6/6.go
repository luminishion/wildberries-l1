/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func way1() {
	ch := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ch
	}()

	ch <- false // using channel
}

func way2() {
	ch := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for range ch {

		}
	}()

	close(ch) // using channel close
}

func way3() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
	}()

	cancel() // using context
}

func main() {
	way1()
	way2()
	way3()

	wg.Wait()
	fmt.Println("all stopped")
}
