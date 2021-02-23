package main

import (
	"bufio"
	"fmt"
	"hackerrank-go/helper"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nd := strings.Split(helper.ReadLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	helper.CheckError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	helper.CheckError(err)
	d := int32(dTemp)

	aTemp := strings.Split(helper.ReadLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		helper.CheckError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	leftRotation(n, d, a)
}

func leftRotation(n, d int32, a []int32) {
	var b = make([]int32, n)
	for i := 0; i < int(n); i++ {
		rotation := i - int(d)
		if rotation < 0 {
			rotation += int(n)
		}
		b[rotation] = a[i]
	}

	fmt.Println(strings.Trim(fmt.Sprint(b), "[]"))
}
