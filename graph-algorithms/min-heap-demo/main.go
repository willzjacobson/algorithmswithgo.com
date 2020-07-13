package main

import (
	minheap "algo/graph-algorithms/min-heap"
	"fmt"
)

func testHeap() {
	heap := minheap.CreateMinHeap()
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
}

func main() {
	testHeap()
}
