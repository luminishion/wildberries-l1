/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import (
	"fmt"
)

func sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	low, high := 0, len(arr)-1
	pivot := (low + high) / 2 // select pivot, could be random / median / first / ...

	arr[pivot], arr[high] = arr[high], arr[pivot] // swap PIVOT and HIGH

	// make everything what is lesser than ARR[HIGH] be below LOW
	// and what is greater or equal be upper LOW
	for i := low; i < high; i++ {
		if arr[i] < arr[high] {
			arr[i], arr[low] = arr[low], arr[i] // swap with LOW
			low++                               // move LOW upper
		}
	}

	arr[low], arr[high] = arr[high], arr[low] // swap LOW and HIGH (formerly pivot)

	// do same for
	sort(arr[:low])   // part what is lesser
	sort(arr[low+1:]) // part what is greater
}

func main() {
	v := []int{1, 5, 8, 4, 2, 9, 4, 6}
	sort(v)
	fmt.Println(v)
}
