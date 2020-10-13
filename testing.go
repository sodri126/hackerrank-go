package main

import "fmt"

func main() {
	input := `RULRD?`
	output := make([]int, 4)
	outputDir := [4][2]int{}
	for _, char := range input {
		switch char {
		case 'R':
			outputDir[0][0]++
			outputDir[1][0]++
			outputDir[2][0]++
			outputDir[3][0]++
		case 'L':
			outputDir[0][0]--
			outputDir[1][0]--
			outputDir[2][0]--
			outputDir[3][0]--
		case 'U':
			outputDir[0][1]++
			outputDir[1][1]++
			outputDir[2][1]++
			outputDir[3][1]++
		case 'D':
			outputDir[0][1]--
			outputDir[1][1]--
			outputDir[2][1]--
			outputDir[3][1]--
		case '?':
			outputDir[0][0]-- // Left
			outputDir[1][0]++ // Right
			outputDir[2][1]++ // Up
			outputDir[3][1]-- // Down
		}
	}

	for _, val := range outputDir {
		// xmin
		if val[0] < output[0] {
			output[0] = val[0]
		}

		// ymin
		if val[1] < output[1] {
			output[1] = val[1]
		}

		// xmax
		if val[0] > output[2] {
			output[2] = val[0]
		}

		// ymax
		if val[1] > output[3] {
			output[3] = val[1]
		}
	}

	fmt.Println(output)
}
