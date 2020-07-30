package main

import (
	"algo/sorting/heap/heap"
	"fmt"
)

// HeapSortInt will sort a list of integers using the heap sort algorithm.
//
// Big O: O(NlogN), where N is the size of the list
func HeapSortInt(nums *[]int) {
	if len(*nums) == 0 {
		return
	}

	h := heap.CreateMaxHeap()
	h.AssignList(nums)

	for i := len(*nums) - 1; i >= 0; i-- {
		(*nums)[i] = h.ExtactMax()
	}
}

func main() {
	// testHeap()
	testHeapSort()
}

func testHeap() {
	heap := heap.CreateMaxHeap()
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
}

func testHeapSort() {
	toSort := []int{}
	HeapSortInt(&toSort)
	fmt.Println(toSort)

	toSort = []int{1}
	HeapSortInt(&toSort)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	HeapSortInt(&toSort)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	HeapSortInt(&toSort)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	HeapSortInt(&toSort)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	HeapSortInt(&toSort)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	HeapSortInt(&toSort)
	fmt.Println(toSort)

	toSort = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	HeapSortInt(&toSort)
	fmt.Println(toSort)
}
