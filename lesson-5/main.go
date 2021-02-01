package main

import (
	"fmt"
	"log"
)

type Node struct {
	next  *Node
	prev  *Node
	value int
}

type List struct {
	count int
	head  *Node
	tail  *Node
}

func (l *List) Print() {
	for tmp, i := l.head, l.count; i > 0; tmp, i = tmp.next, i-1 {
		fmt.Printf("Value %d, Prev: %v, Next: %v \n", tmp.value, tmp.prev, tmp.next)
	}

}

func (l *List) PushBack(node *Node) {
	prev := l.tail
	l.count++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	prev.next = node
	node.prev = prev
	l.tail = node
}

func (l *List) PushFront(node *Node) {

	prev := l.head

	l.count++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = prev
	prev.prev = node
	l.head = node

}

func (l *List) PopBack() {

	if l.count <= 0 {
		log.Fatal("PopBack() called on empty queue")
	}

	del := l.tail
	node := del.prev
	node.next = nil
	del.prev = nil
	l.tail = node
	l.count--
}

func (l *List) PopFront() {
	if l.count <= 0 {
		log.Fatal("PopBack() called on empty queue")
	}

	del := l.head
	head := l.head.next

	head.prev = nil
	del.next = nil
	l.head = head
	l.count--
}

func main() {
	l := &List{}

	node1 := &Node{value: 1}
	node2 := &Node{value: 2}
	node3 := &Node{value: 3}
	node4 := &Node{value: 4}

	fmt.Println("PushBack")
	l.PushBack(node1)
	l.PushBack(node2)
	l.PushBack(node3)

	l.Print()

	fmt.Println("\nPushFront")
	l.PushFront(node4)

	l.Print()
	fmt.Println("\nPopBack")

	l.PopBack()
	l.Print()

	fmt.Println("\nPopFront")

	l.PopFront()
	l.Print()

}
