/*
Удалить i-ый элемент из слайса.
*/

package main

import (
	"fmt"
)

func way1() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 3

	arr = append(arr[:i], arr[i+1:]...)

	fmt.Println(arr)
}

// if order doesn't matter, and no need to remove a lot of items
func way2() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 3

	arr[i] = arr[len(arr)-1] // move last to i
	arr[len(arr)-1] = 0      // set default
	arr = arr[:len(arr)-1]   // truncate

	fmt.Println(arr)
}

func main() {
	way1()
	way2()
}
