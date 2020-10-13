package main

import (
	"bufio"
	"fmt"
	"hackerrank/helper"
	"os"
	"strconv"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

// Complete the printLinkedList function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func printLinkedList(head *SinglyLinkedListNode) {
	for head != nil {
		fmt.Println((*head).data)
		head = head.next
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	llistCount, err := strconv.ParseInt(helper.ReadLine(reader), 10, 64)
	helper.CheckError(err)

	llist := SinglyLinkedList{}
	for i := 0; i < int(llistCount); i++ {
		llistItemTemp, err := strconv.ParseInt(helper.ReadLine(reader), 10, 64)
		helper.CheckError(err)
		llistItem := int32(llistItemTemp)
		llist.insertNodeIntoSinglyLinkedList(llistItem)
	}

	printLinkedList(llist.head)
}
