package main

import "fmt"

// LinkedList : Implementing a linked list using 3 arrays
type LinkedList struct {
	next []int
	key  []int
	prev []int
	head int
	null int
	free int
}

// CreateLinkedList : Creates a LinkedList instance
func CreateLinkedList() *LinkedList {
	null := -1
	return &LinkedList{
		free: null,
		head: null,
		null: null,
	}
}

// Search : return the index containing the specified key
func (l *LinkedList) Search(x int) int {
	n := l.head
	for n != l.null && l.key[n] != x {
		n = l.next[n]
	}

	return n
}

// Insert : insert a new "node"
func (l *LinkedList) Insert(x int) {
	newHead := l.allocateObject()
	if l.head != l.null {
		l.prev[l.head] = newHead
	}
	l.next[newHead] = l.head // new head 'next' points to old head
	l.prev[newHead] = l.null // new head 'prev' is empty by definition
	l.key[newHead] = x       // insert the key
	l.head = newHead         // make this new node the head of the list
}

// allocateObject : find or create a free slot to put a new "node"
func (l *LinkedList) allocateObject() int {
	if l.free == l.null {
		// Either brand new list instance, or no space remaining
		// In either case, create a new index, and return it
		l.next = append(l.next, l.null)
		l.prev = append(l.prev, l.null)
		l.key = append(l.key, l.null)
		return len(l.key) - 1
	}

	// Free slots exist. Move free pointer to next free slot, then return index of the newly allocated slot
	free := l.free
	l.free = l.next[l.free]
	return free
}

// Delete : delete "node" at a particular index
func (l *LinkedList) Delete(x int) {
	if len(l.key) < x+1 {
		panic("list underflow")
	}

	// connect "next" node to the "prev" node to bridge the gap being created
	if l.next[x] != l.null {
		l.prev[l.next[x]] = l.prev[x]
	}
	// connect "prev" node to the "next" node to bridge the gap being created
	if l.prev[x] != l.null {
		l.next[l.prev[x]] = l.next[x]
	}
	// If this node is the head of the list, point the head to the "next" node
	if l.head == x {
		l.head = l.next[x]
	}

	l.freeObject(x)
}

// freeObject : mark an index free
func (l *LinkedList) freeObject(x int) {
	l.key[x] = l.null
	l.prev[x] = l.null
	l.next[x] = l.free
	l.free = x
}

func main() {
	l := CreateLinkedList()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	fmt.Println("List:", l)
	fmt.Println("4 is at index:", l.Search(4))
	fmt.Println("5 is at index:", l.Search(5))
	l.Delete(2)
	l.Delete(1)
	fmt.Println("List:", l)
	l.Insert(5)
	fmt.Println("List:", l)
	l.Insert(6)
	fmt.Println("List:", l)
	l.Insert(7)
	fmt.Println("List:", l)
}
