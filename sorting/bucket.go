package main

import (
	"fmt"
)

// BucketSort will sort a list of integers using the bucket sort algorithm.
//
// Big O: average case is O(n), worst case is O(n^2) (if all data falls in 1 bucket)
func BucketSort(nums []float32) {
	if len(nums) < 2 {
		return
	}

	out := []float32{}
	b := make([][]float32, len(nums))

	for _, v := range nums {
		listInd := int(float32(len(nums)) * v)
		b[listInd] = append(b[listInd], v)
	}
	for i := range b {
		InsertionSortIntBrute(b[i])
		out = append(out, b[i]...)
	}
	for i, v := range out {
		nums[i] = v
	}
}

// InsertionSortIntBrute will sort a list of integers using the insertion sort
// algorithm, using brute force rather than binary search
// to find the index at which to insert each element in the sorted array
//
// Big O (without binary search): O(N^2), where N is the size of the list.
//
// See below for binary search implementation
func InsertionSortIntBrute(list []float32) {
	var sorted []float32
	for _, item := range list {
		sorted = insertInt(sorted, item)
	}

	for i, v := range sorted {
		list[i] = v
	}
}

func insertInt(sorted []float32, item float32) []float32 {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			return append(sorted[:i], append([]float32{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item)
}

func main() {
	toSort := []float32{}
	BucketSort(toSort)
	fmt.Println(toSort)

	toSort = []float32{.1}
	BucketSort(toSort)
	fmt.Println(toSort)

	toSort = []float32{.1, .2, .3}
	BucketSort(toSort)
	fmt.Println(toSort)

	toSort = []float32{.3, .2, .1}
	BucketSort(toSort)
	fmt.Println(toSort)

	toSort = []float32{.5, .3, .5, .3, .5}
	BucketSort(toSort)
	fmt.Println(toSort)

	toSort = []float32{.3, .5, .3, .5, .3}
	BucketSort(toSort)
	fmt.Println(toSort)

	toSort = []float32{.12, .88, .32, .1, .23, .96, .91, .4, .48, .7, .5, .21}
	BucketSort(toSort)
	fmt.Println(toSort)
}
