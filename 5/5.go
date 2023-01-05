/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

package main

import (
	"fmt"
	"time"
)

func receiver(ch chan int) {
	for val := range ch {
		fmt.Println("received", val)
	}
}

func sender(ch chan int) {
	ctr := 0

	for {
		ctr++
		ch <- ctr
		fmt.Println("sending", ctr)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	go receiver(ch)
	go sender(ch)

	time.Sleep(8 * time.Second)
	fmt.Println("exit")
}
