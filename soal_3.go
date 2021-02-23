package main

import "strconv"

func main() {
	number := "10111214"
	var data []int
	for stop, digits := false, 1; !stop; {
		n, _ := strconv.Atoi(number[:digits])
		data = append(data, n)
		stop = true
	}
}
