/*!
 * main.go
 * bucketSortByLinkedList
 *
 * Copyright (c) 2019 Masaharu TASHIRO
 * Released under the MIT license.
 * see https://opensource.org/licenses/MIT
 */

package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Node ...
type Node struct {
	value int
	next  *Node
}

func main() {
	// Parse arguments
	n, err := parseArgs(os.Args)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("N: %d\n", n)

	// Initilize list
	rand.Seed(time.Now().UnixNano())
	list := initList(n)

	// Print unsorted list
	fmt.Println("Unsorted list")
	printList(list)

	// sorting
	sortedList := bucketSort(list, 2*n)

	// Print sorted list
	fmt.Println("Sorted list")
	printList(sortedList)

}

func initList(n int) *Node {
	list := (*Node)(nil)
	for i := 0; i < n; i++ {
		list = &Node{value: rand.Intn(2 * n), next: list}
	}
	return list
}

func bucketSort(list *Node, max int) *Node {

	table := make([]*Node, max+1)
	for node := list; node != nil; {
		nextNode := node.next
		node.next = table[node.value]
		table[node.value] = node
		node = nextNode
	}

	sortedList := &Node{}
	end := sortedList
	for _, nodes := range table {
		for node := nodes; node != nil; node = node.next {
			end.next = node
			end = node
		}
		end.next = nil
	}

	return sortedList.next
}

func printList(list *Node) {
	for node := list; node != nil; node = node.next {
		fmt.Print(node.value, ",")
	}
	fmt.Println()
}

func parseArgs(args []string) (int, error) {

	if len(args) < 2 {
		err := errors.New("Too few arguments")
		return 0, err
	} else if len(args) > 2 {
		err := errors.New("Too much arguments")
		return 0, err
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println(err)
		return 0, err
	}
	if n < 1 {
		err := errors.New("N must be greater than 0")
		return 0, err
	}

	return n, nil
}
