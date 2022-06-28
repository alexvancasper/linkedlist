package main

import (
	"fmt"
	"math/rand"
	LL "path/to/linkedlist"
	"time"
)

type Node struct {
	value int
	next  *Node
}

func mergeTwoLists(l1 *LL.Node, l2 *LL.Node) *LL.Node {
	lst := &LL.Node{}
	p := l1
	q := l2
	curr := lst
	var x, y int = 101, 101
	for p != nil || q != nil {
		if p != nil {
			x = p.Value
		} else {
			x = 101
		}
		if q != nil {
			y = q.Value
		} else {
			y = 101
		}
		if x == y {
			curr.Next = &LL.Node{Value: x}
			n := &LL.Node{Value: y}
			curr = curr.Next
			curr.Next = n
			if p != nil {
				p = p.Next
			}
			if q != nil {
				q = q.Next
			}
		}
		if x < y {
			curr.Next = &LL.Node{Value: x}
			if p != nil {
				p = p.Next
			}
		}
		if x > y {
			curr.Next = &LL.Node{Value: y}
			if q != nil {
				q = q.Next
			}
		}
		curr = curr.Next
	}
	return lst.Next
}

func main() {
	var lst1 LL.LinkedList
	var lst2 LL.LinkedList
	var lst3 LL.LinkedList

	/*
	mergeTwoLists
	*/
	lst1.PushBack(1)
	lst1.PushBack(5)
	lst1.PushBack(8)

	lst2.PushBack(1)
	lst2.PushBack(2)
	lst2.PushBack(4)
	// lst2.PushBack(5)
	lst2.PushBack(6)

	lst1.Print(" ")
	println()
	lst2.Print(" ")

	lst3.Head = mergeTwoLists(lst1.Head, lst2.Head)
	fmt.Println("\nResult:")
	lst3.Print(" ")

	/*
	$ go run main.go
	1 5 8
	1 2 4 6
	Result:
	1 1 2 4 5 6 8
	*/

}