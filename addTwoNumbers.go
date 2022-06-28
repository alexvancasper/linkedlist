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

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	lst := &Node{}
	p := l1
	q := l2
	curr := lst
	carry := 0
	var x, y int = 0, 0
	for p != nil || q != nil {
		if p != nil {
			x = p.value
		} else {
			x = 0
		}
		if q != nil {
			y = q.value
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		curr.next = &Node{value: sum % 10}
		curr = curr.next
		if p != nil {
			p = p.next
		}
		if q != nil {
			q = q.next
		}
	}
	if carry > 0 {
		curr.next = &Node{value: carry}
	}

	return lst.next
}

func main() {
	var lst1 LL.LinkedList
	var lst2 LL.LinkedList

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 100; i++ {
		lst1.PushFront(r1.Intn(10))
	}
	for i := 0; i < 100; i++ {
		lst2.PushFront(r1.Intn(10))
	}

	fmt.Println("number 1: ")
	lst1.Print()
	lst1.ReversePrint()
	fmt.Println("\nnumber 2: ")
	lst2.ReversePrint()

	var lst3 LL.LinkedList
	lst3.head = addTwoNumbers(lst1.head, lst2.head)
	fmt.Println("\nResult of SUM: ")
	lst3.ReversePrint()

}
