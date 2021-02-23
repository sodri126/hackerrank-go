package main

import (
	"fmt"
)

func main() {
	var min, max, res int
	fmt.Scanln(&min, &max)
	for ; min <= max; min++ {
		if rev := reverseNumber(min); rev == min {
			res++
		}
	}
	fmt.Println(res)
}

func reverseNumber(number int) int {
	var res, modN int
	for ; number != 0; number /= 10 {
		modN = number % 10
		res = res*10 + modN
	}
	return res
}
