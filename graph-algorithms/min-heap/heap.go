package minheap

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"fmt"
)

// Application of a min heap where the satellite data is an adjacency-list vertex
// We store the index in the AdjListVertex type itself to be able to use the DecreaseKeys method in our implementation Prim's algorthim
// This may violate the principle of separation of concerns, but *shrug*

type MinHeapNode struct {
	right     *MinHeapNode
	left      *MinHeapNode
	key       float64
	satellite *adjacencylist.AdjListVertex
}

func CreateMinHeapNode(key float64, satellite *adjacencylist.AdjListVertex) *MinHeapNode {
	return &MinHeapNode{
		satellite: satellite,
		key:       key,
	}
}

type MinHeap struct {
	Size  int
	slice []*adjacencylist.AdjListVertex
}

func CreateMinHeap() *MinHeap {
	return &MinHeap{
		slice: []*adjacencylist.AdjListVertex{},
	}
}

// Only need to worry about the nodes that are not leaves
func (h *MinHeap) AssignList(vertices []*adjacencylist.AdjListVertex) {
	if h.Size > 0 {
		panic("This method is only intended for when the heap is empty")
	}
	h.slice = vertices
	h.Size = len(vertices)
	for i := len(h.slice) / 2; i >= 0; i-- {
		h.SendNodeDown(i)
	}
}

func (h *MinHeap) Insert(vertex *adjacencylist.AdjListVertex) {
	vertex.Index = h.Size // assign index to new vertex reflecting the position at the end of the heap
	h.slice = append(h.slice, vertex)
	h.Size++
	h.bringNodeUp(h.Size - 1)
}

func (h *MinHeap) Minimum() *adjacencylist.AdjListVertex {
	if h.Size < 1 {
		panic("heap underflow")
	}
	return h.slice[0]
}

func (h *MinHeap) ExtractMin() *adjacencylist.AdjListVertex {
	if h.Size < 1 {
		panic("heap underflow")
	}
	min := h.slice[0]
	h.slice[0] = h.slice[h.Size-1]
	h.slice[0].Index = 0 // update index to reflect new position in heap
	h.Size--
	h.SendNodeDown(0)
	return min
}

func (h *MinHeap) DecreaseKeys(i int, decreaseTo float64) {
	if h.Size <= i {
		panic("heap underflow")
	}
	h.slice[i].Key = decreaseTo
	h.bringNodeUp(i)
}

func (h *MinHeap) bringNodeUp(i int) {
	parentInd := i / 2
	for h.slice[i].Key < h.slice[parentInd].Key {
		// change index property to match the location swap we're about to do
		h.slice[parentInd].Index = i
		h.slice[i].Index = parentInd
		// swap the node's positions
		h.slice[i], h.slice[parentInd] = h.slice[parentInd], h.slice[i]
		// update parameters and continue looping
		i = parentInd
		parentInd = i / 2
	}
}

func (h *MinHeap) SendNodeDown(i int) {
	smallestInd := i
	smallest := h.slice[i]
	leftInd := i * 2
	rightInd := i*2 + 1

	if h.Size >= leftInd+1 && h.slice[leftInd].Key < smallest.Key {
		smallest = h.slice[leftInd]
		smallestInd = leftInd
	}
	if h.Size >= rightInd+1 && h.slice[rightInd].Key < smallest.Key {
		smallest = h.slice[leftInd]
		smallestInd = rightInd
	}

	if smallestInd != i {
		// change index property to match the location swap we're about to do
		h.slice[smallestInd].Index = i
		h.slice[i].Index = smallestInd
		// swap locations of elements in the heap
		h.slice[smallestInd], h.slice[i] = h.slice[i], h.slice[smallestInd]
		// continue down the heap
		h.SendNodeDown(smallestInd)
	}
}

func TestMinHeap() {
	heap := CreateMinHeap()
	fmt.Println(heap)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   3,
		Value: "a",
	})
	fmt.Println(heap.slice[0].Index)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   2,
		Value: "b",
	})
	fmt.Println(heap.slice[0].Index, heap.slice[1].Index)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   4,
		Value: "c",
	})
	fmt.Println(heap.slice[0].Index, heap.slice[1].Index, heap.slice[2].Index)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   5,
		Value: "d",
	})
	fmt.Println(heap.slice[0].Index, heap.slice[1].Index, heap.slice[2].Index, heap.slice[3].Index)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   6,
		Value: "e",
	})
	fmt.Println(heap)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   7,
		Value: "f",
	})
	fmt.Println(heap)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   8,
		Value: "g",
	})
	fmt.Println(heap)

	heap.Insert(&adjacencylist.AdjListVertex{
		Key:   9,
		Value: "h",
	})
	fmt.Println(heap)

	heap.ExtractMin()
	fmt.Println(heap)

	heap.ExtractMin()
	fmt.Println(heap)

	heap.DecreaseKeys(5, 1)
	fmt.Println(heap)
	heap.DecreaseKeys(5, 1)
	fmt.Println(heap.slice[0].Key, heap.slice[1].Key, heap.slice[2].Key, heap.slice[3].Key, heap.slice[4].Key, heap.slice[5].Key, heap.slice[6].Key, heap.slice[7].Key)
}

func main() {
	TestMinHeap()
}
