package main

import "fmt"

type ListNode struct {
	next *ListNode
	prev *ListNode
	data int
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	null *ListNode
}

func (l *LinkedList) Insert(x int) {
	n := &ListNode{
		data: x,
		next: l.head,
		prev: l.null,
	}

	l.head.prev = n
	l.head = n
}

func (l *LinkedList) Search(x int) *ListNode {
	n := l.head
	for n != l.null && n.data != x {
		n = n.next
	}
	return n
}

func (l *LinkedList) Delete(x int) {
	// find node to delete
	n := l.Search(x)
	// delete it
	n.prev.next = n.next
	n.next.prev = n.prev
}

func CreateLinkedList() *LinkedList {
	// create sentinel to allow for cleaner code
	nullNode := &ListNode{}
	nullNode.next = nullNode
	nullNode.prev = nullNode

	return &LinkedList{
		null: nullNode,
		head: nullNode,
	}
}

func (l *LinkedList) EnumerateList() {
	n := l.head
	for n != l.null {
		fmt.Println("Node data:", n.data)
		n = n.next
	}
}

func main() {
	l := CreateLinkedList()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.EnumerateList()
	fmt.Println("-=-=-")
	l.Delete(3)
	l.Delete(5)
	l.EnumerateList()
	l.Insert(5)
	l.Insert(6)
	fmt.Println("-=-=-")
	l.EnumerateList()
}
