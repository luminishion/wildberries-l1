/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.

Выбрать и обосновать способ завершения работы всех воркеров.
	using context to stop workers
	because if we gonna modify the code later
	then something inside workers may require context to stop too
	so i think it is a more general way than closing a channel
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, jobs chan int, wg *sync.WaitGroup, i int) {
	for {
		select {
		case ctr := <-jobs:
			fmt.Printf("worker %d received %d\n", i, ctr)
		case <-ctx.Done(): // on stop
			wg.Done()
			return
		}
	}
}

func main() {
	fmt.Println("enter number of workers:")

	var n int
	fmt.Scanf("%d", &n)

	jobs := make(chan int)

	//--------------------- start workers ---------------------
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go worker(ctx, jobs, &wg, i)
	}
	//--------------------------------------------------------

	// --------------------- start jobs writer ---------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt) // catch signals

	ctr := 0

	for {
		select {
		case <-time.After(time.Second): // unfreezes every second
			ctr++
			jobs <- ctr

			fmt.Printf("sending %d\n", ctr)
		case <-quit: // on signal to exit
			cancel()
			fmt.Println("shutting down")

			wg.Wait()
			fmt.Println("goroutines stopped")

			return
		}
	}
	//---------------------------------------------------------------

}
