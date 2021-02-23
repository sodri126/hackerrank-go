package helper

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
