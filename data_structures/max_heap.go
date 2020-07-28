package main

import "fmt"

/*
MaxHeap :
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

type MaxHeapNode struct {
	right *MaxHeapNode
	left  *MaxHeapNode
	value int
}

func CreateMaxHeapNode(value int) *MaxHeapNode {
	return &MaxHeapNode{
		value: value,
	}
}

type MaxHeap struct {
	size  int
	slice []int
}

func CreateMaxHeap() *MaxHeap {
	return &MaxHeap{
		slice: []int{},
	}
}

// Only need to worry about the nodes that are not leaves
func (h *MaxHeap) AssignList(nums *[]int) {
	if h.size > 0 {
		panic("This method is only intended for when the heap is empty")
	}
	h.slice = *nums
	h.size = len(*nums)
	for i := (len(h.slice) / 2) - 1; i >= 0; i-- {
		h.SendNodeDown(i)
	}
}

func (h *MaxHeap) Insert(val int) {
	h.slice = append(h.slice, val)
	h.size++
	h.bringNodeUp(h.size - 1)
}

func (h *MaxHeap) Maximum() int {
	if h.size < 1 {
		panic("heap underflow")
	}
	return h.slice[0]
}

func (h *MaxHeap) ExtactMax() int {
	if h.size < 1 {
		panic("heap underflow")
	}
	max := h.slice[0]
	h.slice[0] = h.slice[h.size-1]
	h.size--
	h.SendNodeDown(0)
	return max
}

func (h *MaxHeap) IncreaseKeys(i int, increaseTo int) {
	if h.size <= i {
		panic("heap underflow")
	}
	h.slice[i] = increaseTo
	h.bringNodeUp(i)
}

func (h *MaxHeap) bringNodeUp(i int) {
	parentInd := (i - 1) / 2
	for h.slice[i] > h.slice[parentInd] {
		h.slice[i], h.slice[parentInd] = h.slice[parentInd], h.slice[i]
		i = parentInd
		parentInd = (i - 1) / 2
	}
}

func (h *MaxHeap) SendNodeDown(i int) {
	largestInd := i
	largest := h.slice[i]
	leftInd := (i+1)*2 - 1  // accounting for off-by-1 error due to Go starting indices at 0 (if Go started at 1, would be 2i)
	rightInd := (i + 1) * 2 // accounting for off-by-1 error due to Go starting indices at 0 (if Go started at 1, would be 2i+1)

	if h.size >= leftInd+1 && h.slice[leftInd] > largest {
		largest = h.slice[leftInd]
		largestInd = leftInd
	}
	if h.size >= rightInd+1 && h.slice[rightInd] > largest {
		largest = h.slice[rightInd]
		largestInd = rightInd
	}

	if largestInd != i {
		h.slice[largestInd], h.slice[i] = h.slice[i], h.slice[largestInd]
		h.SendNodeDown(largestInd) // subtract 1 to pass in true index of slice
	}
}

func testHeap() {
	heap := CreateMaxHeap()
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

	heap.ExtactMax()
	fmt.Println(heap)

	heap.ExtactMax()
	fmt.Println(heap)

	heap.IncreaseKeys(5, 12)
	fmt.Println(heap)
	heap.IncreaseKeys(5, 12)
	fmt.Println(heap)

	fmt.Println("-=-=-")
	heap2 := CreateMaxHeap()
	heap2.AssignList(&[]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(heap2)
}

func main() {
	testHeap()
}
