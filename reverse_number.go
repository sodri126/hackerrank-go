package main

import "fmt"

func reverse_number(number int) (res int) {
	for number != 0 {
		res *= 10
		temp := number % 10
		res += temp
		number = (number - temp) / 10
	}
	return
}

func main() {
	x := 123456789
	fmt.Println(reverse_number(x))
}
