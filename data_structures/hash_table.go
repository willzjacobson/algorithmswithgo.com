package main

import (
	"fmt"
	"math"
)

type HashTable struct {
	slots  []*LaLinkedList
	length int
}

func (t *HashTable) hashDivide(x int) int {
	return x % t.length
}

func (t *HashTable) hashMultiply(x int) int {
	A := (math.Sqrt(5) - 1) / 2
	_, dec := math.Modf(A * float64(x))
	return int(float64(t.length) * dec)
}

func CreateHashTable(numSlots int) *HashTable {
	slots := make([]*LaLinkedList, numSlots)
	for i := 0; i < numSlots; i++ {
		slots[i] = CreateLaLinkedList()
	}

	return &HashTable{
		length: numSlots,
		slots:  slots,
	}
}

func (t *HashTable) Insert(x int) {
	slot := t.hashMultiply(x)
	t.slots[slot].Insert(x)
}

func (t *HashTable) Search(x int) *LaListNode {
	slot := t.hashMultiply(x)
	return t.slots[slot].Search(x)
}

func (t *HashTable) Delete(x int) {
	slot := t.hashMultiply(x)
	t.slots[slot].Delete(x)
}

func main() {
	t := CreateHashTable(8)
	t.Insert(6)
	t.slots[5].EnumerateList()
	t.slots[6].EnumerateList()
	fmt.Println("-=-=-")
	t.Insert(14)
	t.slots[5].EnumerateList()
	t.slots[6].EnumerateList()
	fmt.Println("-=-=-")
	fmt.Println(t.Search(6).data) // found
	fmt.Println(t.Search(7).data) // not found; value is that of the null sentinel in the linked list in the slot
	fmt.Println("-=-=-")
	t.Delete(6) // delete node
	t.Delete(7) // node doesn't exist; does nothing
	t.slots[6].EnumerateList()
	t.slots[5].EnumerateList()
}

// Linked List implementation
type LaListNode struct {
	next *LaListNode
	prev *LaListNode
	data int
}

type LaLinkedList struct {
	head *LaListNode
	tail *LaListNode
	null *LaListNode
}

func (l *LaLinkedList) Insert(x int) {
	n := &LaListNode{
		data: x,
		next: l.head,
		prev: l.null,
	}

	l.head.prev = n
	l.head = n
}

func (l *LaLinkedList) Search(x int) *LaListNode {
	n := l.head
	for n != l.null && n.data != x {
		n = n.next
	}
	return n
}

func (l *LaLinkedList) Delete(x int) {
	// find node to delete
	n := l.Search(x)
	// delete it
	n.prev.next = n.next
	n.next.prev = n.prev
}

func CreateLaLinkedList() *LaLinkedList {
	// create sentinel to allow for cleaner code
	nullNode := &LaListNode{}
	nullNode.next = nullNode
	nullNode.prev = nullNode

	return &LaLinkedList{
		null: nullNode,
		head: nullNode,
	}
}

func (l *LaLinkedList) EnumerateList() {
	n := l.head
	for n != l.null {
		fmt.Println("Node data:", n.data)
		n = n.next
	}
}
