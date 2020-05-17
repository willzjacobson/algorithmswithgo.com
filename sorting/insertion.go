package main

import "fmt"

// InsertionSortInt will sort a list of integers using the insertion sort algorithm
// using binary search to find the index at which to insert each element
//
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	var sorted []int
	for _, item := range list {
		i := findInsertionIndex(sorted, item)
		sorted = append(sorted[:i], append([]int{item}, sorted[i:]...)...)
	}

	for i, v := range sorted {
		list[i] = v
	}
}

// uses binary search to find the index at which to insert this item
func findInsertionIndex(list []int, item int) int {
	if len(list) == 0 {
		return 0
	}

	high := len(list) - 1
	low := 0
	mid := len(list) / 2

	for {
		if item < list[mid] {
			if mid == 0 || item >= list[mid-1] {
				return mid
			} else if mid-low == 1 {
				mid = low
			} else {
				high, mid = mid, mid-((mid-low)/2)
			}
		} else if item > list[mid] {
			if mid == len(list)-1 || item <= list[mid+1] {
				return mid + 1
			} else if high-mid == 1 {
				mid = high
			} else {
				low, mid = mid, mid+((high-mid)/2)
			}
		} else {
			return mid
		}
	}
}

func main() {
	toSort := []int{}
	InsertionSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{1}
	InsertionSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	InsertionSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	InsertionSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	InsertionSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	InsertionSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	InsertionSortInt(toSort)
	fmt.Println(toSort)
}
