package main

import "fmt"

func main() {
	data := []int{10, 2, 3, 5, 19}
	sum := 5

	fmt.Println(HasRepairedSum(data, sum))
}

func HasRepairedSum(data []int, sum int) bool {
	tmp := make(map[int]bool, 0)
	for _, d := range data {
		if tmp[d] == true {
			return true
		}
		tmp[sum-d] = true
	}
	return false
}
