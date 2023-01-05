/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

func sleep2(d time.Duration) {
	wait := time.Now().Add(d)
	for wait.After(time.Now()) {

	}
}

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println(1)
	sleep(time.Millisecond * 2)
	fmt.Println(2)
}
