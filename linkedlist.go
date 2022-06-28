package LinkedList

import "fmt"

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
	len  int
}

func (l *LinkedList) GetSize() int {
	return l.len
}

func (l *LinkedList) GetByNum(elem int) int {
	ptr := l.head

	for i := 0; i < l.len; i++ {
		if elem == i {
			return ptr.value
		}
		ptr = ptr.next
	}
	return -1
}

func (l *LinkedList) PushBack(val int) {
	n := node{}
	n.value = val

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) PushFront(val int) {
	newHead := node{value: val, next: l.head}
	l.head = &newHead
	l.len++
}

func (l *LinkedList) PushAt(val int, position int) {
	n := node{value: val}
	ptr := l.head
	ctr := 0
	if position == 0 {
		l.PushFront(val)
	}
	if position == l.len+1 {
		l.PushBack(val)
	}
	if position < 0 || position > l.len {
		return
	}
	for ptr != nil {
		if ctr+1 == position {
			n.next = ptr.next
			ptr.next = &n
			break
		}
		ctr++
		if ptr != nil {
			ptr = ptr.next
		}
	}
	l.len++
}

func (l *LinkedList) Print() {
	ptr := l.head
	for ptr != nil {
		fmt.Printf("%d", ptr.value)
		if ptr.next == nil {
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) ReversePrint() {
	var lst LinkedList
	ptr := l.head
	for ptr != nil {
		lst.PushFront(ptr.value)
		if ptr.next == nil {
			break
		}
		ptr = ptr.next
	}
	lst.Print()
}
