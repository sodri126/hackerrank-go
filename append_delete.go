package main

import "fmt"

func main() {
	var s, t string
	var k, totalOperation int
	fmt.Scanf("%s\n%s\n%d", &s, &t, &k)
	totalS, totalT := len(s), len(t)
	switch {
	case s == t:
		fmt.Println("Yes")
	case s != t:
		if totalS == 0 {
			totalOperation = totalT
			printYesOrNo(totalOperation, k, true)
			return
		}

		if totalT == 0 {
			totalOperation = totalS
			printYesOrNo(totalOperation, k, true)
			return
		}

		forLoop := 0
		if totalS > totalT {
			forLoop = totalT
			totalOperation = totalS - totalT
		} else {
			forLoop = totalS
			totalOperation = totalT - totalS
		}
		tempIndex := -1
		for i := 0; i < forLoop; i++ {
			//fmt.Printf("s[%d]: %c, t[%d]: %c\n", i, s[i], i, t[i])
			if s[i] != t[i] {
				tempIndex = i
				break
				//totalOperation += 2
			}
		}
		if tempIndex > -1 {
			totalOperation += (forLoop - tempIndex) * 2
		}
		printYesOrNo(totalOperation, k, totalS >= totalT)
	}
}

func printYesOrNo(totalOperation, k int, isTrue bool) {
	fmt.Println("Total Operation:", totalOperation, " K: ", k)
	if totalOperation <= k && isTrue || totalOperation%2 == k%2 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
