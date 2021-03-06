/*!
 * main.go
 * shakerSortByDoublyLinkedList
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
	prev  *Node
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

	for i := 0; i < n*n/10; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	// sorting
	sortedList := shakerSort(list)

	for i := 0; i < n*n/10; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	// Print sorted list
	fmt.Println("Sorted list")
	printList(sortedList)

}

func initList(n int) *Node {
	list := &Node{}
	end := list
	for i := 0; i < n; i++ {
		end.next = &Node{value: rand.Intn(2 * n), next: nil, prev: end}
		end = end.next
	}
	list = list.next
	list.prev = nil
	return list
}

func shakerSort(list *Node) *Node {
	node := list
	begin := (*Node)(nil)
	end := (*Node)(nil)

	for {
		isSorted := true
		for ; node.next != end; node = node.next {
			if node.value > node.next.value {
				isSorted = false
				swap(&node.value, &node.next.value)
			}
		}
		printList(list)
		if isSorted == true {
			break
		}
		end = node
		for ; node.prev != begin; node = node.prev {
			if node.prev.value > node.value {
				isSorted = false
				swap(&node.value, &node.prev.value)
			}
		}
		begin = node
		printList(list)
		if isSorted == true {
			break
		}
	}
	return list
}

func swap(x *int, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
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
