package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, m int

	fmt.Scanf("%d\n", &n)
	arr1 := make(map[int]int, 0)
	var j int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &j)
		arr1[j]++
	}

	fmt.Scanf("%d\n", &m)
	arr2 := make(map[int]int, 0)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &j)
		arr2[j]++
	}

	var arr3 []int
	for index, _ := range arr2 {
		if _, ok := arr1[index]; !ok || arr1[index] != arr2[index] {
			arr3 = append(arr3, index)
		}
	}

	sort.Sort(sort.IntSlice(arr3))
	fmt.Println(strings.Trim(fmt.Sprint(arr3), "[]"))

	strings.ToTitle()
}
