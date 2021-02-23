package main

import (
	"fmt"
)

func main() {
	var b, k, d, r int
	fmt.Scan(&b, &k, &d)
	arrK, arrD := make([]int, k), make([]int, d)

	for i := 0; i < k; i++ {
		fmt.Scan(&arrK[i])
	}
	for i := 0; i < d; i++ {
		fmt.Scan(&arrD[i])
	}

	r = -1
	// complexity O(m*n)
	for i := 0; i < k; i++ {
		for j := 0; j < d; j++ {
			if arrK[i]+arrD[j] <= b && arrK[i]+arrD[j] >= r {
				r = arrK[i] + arrD[j]
			}
		}
	}

	// complexity O(n log(n))

	fmt.Println(r)
}
