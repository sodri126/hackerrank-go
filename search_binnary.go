package main

import (
	"fmt"
)

func binary_search(arr []int, search int) (index int) {
	index = -1
	for high, low := len(arr), 0; low <= high; {
		mid := (low + high) / 2
		if mid >= len(arr) {
			return
		}

		if arr[mid] == search {
			index = mid
			return
		}

		if arr[mid] > search {
			high = mid - 1
			continue
		}

		if arr[mid] < search {
			low = mid + 1
			continue
		}
	}
	return
}
func main() {
	arr := []int{1, 2, 3, 20, 30, 50}
	search := 50
	fmt.Println("index 30: ", binary_search(arr, search))
}
