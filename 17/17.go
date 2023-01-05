/*
Реализовать бинарный поиск встроенными методами языка.
*/

package main

import (
	"fmt"
)

func search(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}

	mid := len(arr) / 2

	// if we should search in left part
	if arr[mid] > val {
		return search(arr[:mid], val)
	}

	// if we should search in right part
	if arr[mid] < val {
		ret := search(arr[mid+1:], val)

		// for the -1 return
		if ret >= 0 {
			ret += mid + 1
		}

		return ret
	}

	return mid
}

func main() {
	items := []int{1, 2, 3, 30, 41, 599, 987, 8765, 67757846}
	fmt.Println(search(items, 8765))
}
