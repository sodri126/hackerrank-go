package main

import (
	"bufio"
	"fmt"
	"hackerrank/helper"
	"os"
	"strconv"
	"strings"
)

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	var minVal int32 = 0
	magicSquare := [][][]int32{
		{
			{8, 1, 6},
			{3, 5, 7},
			{4, 9, 2},
		},
		{
			{4, 3, 8},
			{9, 5, 1},
			{2, 7, 6},
		},
		{
			{2, 9, 4},
			{7, 5, 3},
			{6, 1, 8},
		},
		{
			{6, 7, 2},
			{1, 5, 9},
			{8, 3, 4},
		},
		{
			{6, 1, 8},
			{7, 5, 3},
			{2, 9, 4},
		},
		{
			{8, 3, 4},
			{1, 5, 9},
			{6, 7, 2},
		},
		{
			{4, 9, 2},
			{3, 5, 7},
			{8, 1, 6},
		},
		{
			{2, 7, 6},
			{9, 5, 1},
			{4, 3, 8},
		},
	}

	for index, val := range magicSquare {
		var totalForming int32 = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				totalForming += Abs(val[i][j] - s[i][j])
			}
		}

		if index == 0 {
			minVal = totalForming
			continue
		}

		if totalForming < minVal {
			minVal = totalForming
		}
	}
	return minVal
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	helper.CheckError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(helper.ReadLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			helper.CheckError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
