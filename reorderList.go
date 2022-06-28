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

func reorderList(head *Node) {
	p := head
	var q *Node
	ctr := 0
	var lastNode *Node
	for p != nil {

		if ctr%2 != 0 {
			q = p
			for q != nil {
				if q.Next == nil {
					break
				}
				if q.Next.Next == nil {
					lastNode = q.Next
					q.Next = nil
					lastNode.Next = p.Next
					p.Next = lastNode
					break
				}
				if q != nil {
					q = q.Next
				}
			}
		}
		if ctr != 0 {
			if p != nil {
				p = p.Next
			}
		}
		ctr++
	}
}

func main() {
	var lst1 LL.LinkedList
	/*
		reorderList
	*/
	lst1.PushBack(1)
	lst1.PushBack(2)
	lst1.PushBack(3)
	lst1.PushBack(4)
	lst1.PushBack(5)
	lst1.Print(" ")
	reorderList(lst1.Head)
	fmt.Println("\nReordered")
	lst1.Print(" ")

	/*
	$ go run main.go
	1 2 3 4 5
	Reordered
	1 5 2 4 3
	*/

}