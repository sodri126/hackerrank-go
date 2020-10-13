package main

import (
	"bufio"
	"fmt"
	"hackerrank/helper"
	"os"
	"strconv"
)

// Complete the matchingStrings function below.
func matchingStrings(strings []string, queries []string) []int32 {
	var resultMatch = make([]int32, len(queries))
	for i := 0; i < len(queries); i++ {
		for _, str := range strings {
			if queries[i] == str {
				resultMatch[i] += 1
			}
		}
	}
	return resultMatch
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	helper.CheckError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	stringsCount, err := strconv.ParseInt(helper.ReadLine(reader), 10, 64)
	helper.CheckError(err)

	var strings []string

	for i := 0; i < int(stringsCount); i++ {
		stringsItem := helper.ReadLine(reader)
		strings = append(strings, stringsItem)
	}

	queriesCount, err := strconv.ParseInt(helper.ReadLine(reader), 10, 64)
	helper.CheckError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := helper.ReadLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(strings, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}
