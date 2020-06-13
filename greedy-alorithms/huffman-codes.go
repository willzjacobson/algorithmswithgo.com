package main

import "fmt"

/*
Encoding a text file is a great way to reduce its size. You can use variable length codes, and fixed length codes. Prefix codes are variable length codes where no code starts with any other code in the set (this makes it simpler to decode). Variable length/prefix codes often perform better than fixed length codes, bc the average length will be less with variable length codes if you make them as short as they can be.

Huffman created an algorthm for assigning a certain system of prefix codes (Huffman codes) to text. The inputs are the set of characters and the frequency with which each character occurs in the text. The algorithm uses a min heap, and a  basic binary tree implementation (not a binary search tree). The algorthm is greedy in that it continually merges nodes pertaining to the 2 characters with the *lowest* frequency. Thus, the lower frequency characters wind up at the bottom of the tree. Since the depth of a node == the length of its resulting code in bits, the greediness results in the lowest size for the encoded text.
*/

// Huffman : Create a tree of Huffman codes for each character appearing in a text file
// O(nlogn) since the loop is O(n) and each heap operation is logn
func Huffman(chars []TreeNode) *TreeNode {
	heap := CreateMinHeap()
	heap.AssignList(&chars)

	for i := 0; i < len(chars)-1; i++ {
		fmt.Println("-=-=-")
		heap.Enumerate() // useful logging

		// create new tree node to merge 2 pieces of character info
		node := TreeNode{}
		node.left = heap.ExtactMin()
		node.right = heap.ExtactMin()
		node.freq = node.left.freq + node.right.freq // freq of new node is sum of the 2 characters
		heap.Insert(node)                            // insert the new node back into the min heap (the priority queue)
		heap.Enumerate()                             // useful logging
	}

	return heap.ExtactMin()
}

func main() {
	chars := []TreeNode{
		{char: "a", freq: 45},
		{char: "b", freq: 13},
		{char: "c", freq: 12},
		{char: "d", freq: 16},
		{char: "e", freq: 9},
		{char: "f", freq: 5},
	}
	tree := Huffman(chars)
	fmt.Println("Huffman", tree)
}

// -=-=-=- Tree Implementation

type TreeNode struct {
	char  string
	freq  int
	left  *TreeNode
	right *TreeNode
}

func CreateTreeNode(freq int) *TreeNode {
	return &TreeNode{
		freq: freq,
	}
}

// -=-=-=- Min Heap implementation

type MinHeapNode struct {
	right *MinHeapNode
	left  *MinHeapNode
	value TreeNode
}

func CreateMinHeapNode(value TreeNode) *MinHeapNode {
	return &MinHeapNode{
		value: value,
	}
}

type MinHeap struct {
	size  int
	slice []TreeNode
}

func CreateMinHeap() *MinHeap {
	return &MinHeap{
		slice: []TreeNode{},
	}
}

func (h *MinHeap) Enumerate() {
	for i := 0; i < h.size; i++ {
		fmt.Print(h.slice[i].freq, ",")
	}
	fmt.Println("")
}

// Only need to worry about the nodes that are not leaves
func (h *MinHeap) AssignList(chars *[]TreeNode) {
	if h.size > 0 {
		panic("This method is only intended for when the heap is empty")
	}
	h.slice = *chars
	h.size = len(*chars)
	for i := len(h.slice) / 2; i >= 0; i-- {
		h.SendNodeDown(i)
	}
}

func (h *MinHeap) Insert(node TreeNode) {
	if len(h.slice) <= h.size {
		// must make slice longer
		h.slice = append(h.slice, node)
	} else {
		// overwrite old value that is not in use
		h.slice[h.size] = node
	}

	h.size++
	h.bringNodeUp(h.size - 1)
}

func (h *MinHeap) Minimum() TreeNode {
	if h.size < 1 {
		panic("heap underflow")
	}
	return h.slice[0]
}

func (h *MinHeap) ExtactMin() *TreeNode {
	// fmt.Println("EXTRACTING FROM:", h.slice)
	if h.size < 1 {
		panic("heap underflow")
	}
	min := h.slice[0]
	h.slice[0] = h.slice[h.size-1]
	h.size--
	h.SendNodeDown(0)
	return &min
}

func (h *MinHeap) bringNodeUp(i int) {
	parentInd := i / 2
	for h.slice[i].freq < h.slice[parentInd].freq {
		h.slice[i], h.slice[parentInd] = h.slice[parentInd], h.slice[i]
		i = parentInd
		parentInd = i / 2
	}
}

func (h *MinHeap) SendNodeDown(i int) {
	smallestInd := i
	smallest := h.slice[i]
	leftInd := i * 2
	rightInd := i*2 + 1

	if h.size >= leftInd+1 && h.slice[leftInd].freq < smallest.freq {
		smallest = h.slice[leftInd]
		smallestInd = leftInd
	}
	if h.size >= rightInd+1 && h.slice[rightInd].freq < smallest.freq {
		smallest = h.slice[leftInd]
		smallestInd = rightInd
	}

	if smallestInd != i {
		h.slice[smallestInd], h.slice[i] = h.slice[i], h.slice[smallestInd]
		h.SendNodeDown(smallestInd)
	}
}
