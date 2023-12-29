package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) PushBack(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) PushFront(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	l.head = newNode
	newNode.next = temp

}

func (l *List) InsertAfter(valueAfter int, value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.data == valueAfter {
			temp := curr.next
			curr.next = newNode
			newNode.next = temp
			return
		}
		curr = curr.next
	}

}

func (l *List) InsertBefore(valueBefor int, value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	if l.head.data == valueBefor {
		newNode := &Node{data: value, next: l.head}
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.next.data == valueBefor {
			newNode.next = curr.next
			curr.next = newNode
			return
		}
		curr = curr.next
	}

}

func (l *List) MoveAfter(valueAfter int, value int) {
	newNode := &Node{data: value}
	l.remove(value)

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.data == valueAfter {
			temp := curr.next
			curr.next = newNode
			newNode.next = temp
			return
		}
		curr = curr.next
	}

}

func (l *List) MoveBefore(valueBefor int, value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	if l.head.data == valueBefor {
		newNode := &Node{data: value, next: l.head}
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.next.data == valueBefor {
			l.remove(value)
			temp := curr.next
			curr.next = newNode
			newNode.next = temp
			return
		}
		curr = curr.next
	}

}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func main() {
	list := &List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.InsertAfter(2, 3)
	printList(list)
	list.InsertBefore(1, 4)
	printList(list)
	list.PushFront(0)
	printList(list)
	list.MoveAfter(2, 1)
	printList(list)
	list.MoveBefore(1, 3)
	printList(list)
	list.remove(2)

	fmt.Println("List: ")
	printList(list)
}

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
