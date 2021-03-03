package main

import "fmt"

func main() {
	a := 1
	b := 2
	// want to swap value a=2, b=1
	a, b = b, a
	fmt.Println("A: ", a, "B: ", b)

	a = -1
	b = 2
	// want to swap value a=2, b=1 with calculate math
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("A: ", a, "B: ", b)
}
