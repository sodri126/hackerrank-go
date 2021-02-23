package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type MinimumLoss struct {
	Year  int
	Value int64
}

func main() {
	var total int
	var losses []MinimumLoss
	reader := bufio.NewReaderSize(os.Stdin, 1*1)
	fmt.Fscanln(reader, &total)
	var val int64
	for i := 0; i < total; i++ {
		fmt.Fscan(reader, &val)
		minimumLoss := MinimumLoss{
			Year:  i,
			Value: val,
		}
		losses = append(losses, minimumLoss)
	}

	sort.SliceStable(losses, func(i, j int) bool {
		return losses[i].Value > losses[j].Value
	})

	var minimum int64 = math.MaxInt64
	for i := 0; i < len(losses)-1; i++ {
		if losses[i+1].Year > losses[i].Year && losses[i].Value-losses[i+1].Value < minimum {
			minimum = losses[i].Value - losses[i+1].Value
		}
	}

	fmt.Println(minimum)
}
