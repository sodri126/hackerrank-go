package main

import (
	"fmt"
	"os"
)

func main() {
	var total, max int
	_, err := fmt.Scanln(&total)
	if err != nil {
		fmt.Println(err, os.Stderr)
		return
	}
	arr := make([]int, 100)
	for i := 0; i < total; i++ {
		var input int
		_, err = fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error: ", err, os.Stderr)
			return
		}
		arr[input]++
	}
	for i := 0; i < total-1; i++ {
		if arr[i]+arr[i+1] > max {
			max = arr[i] + arr[i+1]
		}
	}

	fmt.Println(max)
}
