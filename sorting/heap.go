package main

import "fmt"

// HeapSortInt will sort a list of integers using the heap sort algorithm.
//
// Big O: O(NlogN), where N is the size of the list
func HeapSortInt(list []int) {
	fmt.Println("Run Heap Sort")
}

func main() {
	toSort := []int{}
	HeapSortInt(toSort)
	fmt.Println(toSort)

	// toSort = []int{1}
	// HeapSortInt(toSort)
	// fmt.Println(toSort)

	// toSort = []int{1, 2, 3}
	// HeapSortInt(toSort)
	// fmt.Println(toSort)

	// toSort = []int{3, 2, 1}
	// HeapSortInt(toSort)
	// fmt.Println(toSort)

	// toSort = []int{5, 3, 5, 3, 5}
	// HeapSortInt(toSort)
	// fmt.Println(toSort)

	// toSort = []int{3, 5, 3, 5, 3}
	// HeapSortInt(toSort)
	// fmt.Println(toSort)

	// toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	// HeapSortInt(toSort)
	// fmt.Println(toSort)
}
