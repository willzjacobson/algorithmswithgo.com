package main

import (
	"fmt"
)

/*
MinHeap :
A heap is represented by an array A that has 2 attributes:
	A.length, which gives the number of elements in the array
	A.heapSize, which represents how many elements in the heap are stored within A
While A[1...A.length] may contain numbers,
only the elements A[1...A.heapSize], where 1 <= A.heapSize <= A.length are valid members of the heap.

The root of the tree is A[1]
We can compute the indices of an element's parent, left, and right child (in this example we index at 1, not 0):
	parent element is i/2
	left child is 2i
	right child is 2i+1

The subarray A[(n/2)+1...n] are all leaves of the tree (and thus are each a 1-element heap)
Thus, when converting an array to a heap, we need only address the elements 1...n/2

Basic heap operations run in time with the height of the tree representation: O(log n)
*/

type MinHeapNode struct {
	right *MinHeapNode
	left  *MinHeapNode
	value int
}

func CreateMinHeapNode(value int) *MinHeapNode {
	return &MinHeapNode{
		value: value,
	}
}

type MinHeap struct {
	size  int
	slice []int
}

func CreateMinHeap() *MinHeap {
	return &MinHeap{
		slice: []int{},
	}
}

// Only need to worry about the nodes that are not leaves
func (h *MinHeap) AssignList(nums *[]int) {
	if h.size > 0 {
		panic("This method is only intended for when the heap is empty")
	}
	h.slice = *nums
	h.size = len(*nums)
	for i := (len(h.slice) - 1) / 2; i >= 0; i-- {
		h.SendNodeDown(i)
	}
}

func (h *MinHeap) Insert(val int) {
	h.slice = append(h.slice, val)
	h.size++
	h.bringNodeUp(h.size - 1)
}

func (h *MinHeap) Minimum() int {
	if h.size < 1 {
		panic("heap underflow")
	}
	return h.slice[0]
}

func (h *MinHeap) ExtactMin() int {
	if h.size < 1 {
		panic("heap underflow")
	}
	min := h.slice[0]
	h.slice[0] = h.slice[h.size-1]
	h.size--
	h.SendNodeDown(0)
	return min
}

func (h *MinHeap) DecreaseKeys(i int, decreaseTo int) {
	if h.size <= i {
		panic("heap underflow")
	}
	h.slice[i] = decreaseTo
	h.bringNodeUp(i)
}

func (h *MinHeap) bringNodeUp(i int) {
	parentInd := (i - 1) / 2
	for h.slice[i] < h.slice[parentInd] {
		h.slice[i], h.slice[parentInd] = h.slice[parentInd], h.slice[i]
		i = parentInd
		parentInd = (i - 1) / 2
	}
}

func (h *MinHeap) SendNodeDown(i int) {
	smallestInd := i
	smallest := h.slice[i]
	leftInd := (i+1)*2 - 1  // accounting for off-by-1 error due to Go starting indices at 0 (if Go started at 1, would be 2i)
	rightInd := (i + 1) * 2 // accounting for off-by-1 error due to Go starting indices at 0 (if Go started at 1, would be 2i+1)

	if h.size >= leftInd+1 && h.slice[leftInd] < smallest {
		smallest = h.slice[leftInd]
		smallestInd = leftInd
	}
	if h.size >= rightInd+1 && h.slice[rightInd] < smallest {
		smallest = h.slice[rightInd]
		smallestInd = rightInd
	}

	if smallestInd != i {
		h.slice[smallestInd], h.slice[i] = h.slice[i], h.slice[smallestInd]
		h.SendNodeDown(smallestInd)
	}
}

func testHeap() {
	heap := CreateMinHeap()
	fmt.Println(heap)

	heap.Insert(3)
	fmt.Println(heap)

	heap.Insert(2)
	fmt.Println(heap)

	heap.Insert(4)
	fmt.Println(heap)

	heap.Insert(5)
	fmt.Println(heap)

	heap.Insert(6)
	fmt.Println(heap)

	heap.Insert(7)
	fmt.Println(heap)

	heap.Insert(8)
	fmt.Println(heap)

	heap.Insert(9)
	fmt.Println(heap)

	heap.ExtactMin()
	fmt.Println(heap)

	heap.ExtactMin()
	fmt.Println(heap)

	heap.DecreaseKeys(5, 1)
	fmt.Println(heap)
	heap.DecreaseKeys(5, 1)
	fmt.Println(heap)

	fmt.Println("-=-=-")
	heap2 := CreateMinHeap()
	heap2.AssignList(&[]int{7, 6, 5, 4, 3, 2, 1})
	fmt.Println(heap2)
}

func main() {
	testHeap()
}
