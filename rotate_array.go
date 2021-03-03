package main

import "fmt"

func rotate(arr []int, totalRotated int, isLeft bool) {
	for i := 0; i < len(arr)-totalRotated; i++ {
		arr[i+totalRotated] = arr[i+totalRotated] + arr[i]
		arr[i] = arr[i+totalRotated] - arr[i]
		arr[i+totalRotated] = arr[i+totalRotated] - arr[i]
		//arr[i+totalRotated], arr[i] = arr[i], arr[i+totalRotated]
	}
}

func rotate1(arr []int, totalRotated int, isLeft bool) []int {
	if isLeft {
		arr = append(arr[totalRotated:], arr[:totalRotated]...)
	} else {
		arr = append(arr[len(arr)-totalRotated:], arr[:len(arr)-totalRotated]...)
	}
	return arr
}

func main() {
	arr := []int{5, 2, 3, 6, 1, 9, 10, 2}
	totalRotated := 3
	fmt.Println("Before rotated: ", arr)
	fmt.Println("Rotated left: ", totalRotated)
	//rotate(arr, totalRotated, true)
	arr = rotate1(arr, totalRotated, false)
	fmt.Println("After rotated: ", arr)
}
