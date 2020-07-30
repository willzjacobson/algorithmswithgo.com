package main

import (
	"fmt"
	"math/rand"
)

// Prefered solution
func quickSort(nums []int, start, end int) {
	if start < end {
		// randomize which element we use as a pivot
		// to prevent n^2 runtimes for sorted or inversely sorted inputs
		p := rand.Intn(end-start) + start
		nums[end], nums[p] = nums[p], nums[end]

		// all values in 'nums' < nums[pivot] go to the left of pivot,
		// all values in 'nums' > nums[pivot]go to the right of pivot.
		pivot := partition(nums, start, end)

		// Sort subarrays to the left and right of pivot
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	comp := nums[end]
	pivot := start - 1
	// Loop through to partially order the subarray, and determine where pivot should be
	for i := start; i < end; i++ {
		if nums[i] < comp {
			nums[i], nums[pivot+1] = nums[pivot+1], nums[i]
			pivot++
		}
	}

	// Put pivot element in its correct spot
	nums[end], nums[pivot+1] = nums[pivot+1], nums[end]
	return pivot + 1
}

// -=-=-=-=-=-

// This solution is NOT 'in place'
// The 2nd and 3rd params aren't actually used,
// it was just convenient for the 2 quicksort implementations to have the same signature for testing
func quickSortSimple(nums []int, start, end int) {
	if len(nums) <= 1 {
		return
	}

	var higher []int
	var lower []int

	midInd := len(nums) / 2
	mid := nums[midInd]

	for i, v := range nums {
		if i == midInd {
			continue
		}
		if v < mid {
			lower = append(lower, v)
		} else {
			higher = append(higher, v)
		}
	}
	quickSortSimple(lower, start, end)
	quickSortSimple(higher, start, end)

	// compose fully sorted array
	sorted := append(lower, append([]int{mid}, higher...)...)
	// overlay sorted order onto original array
	for i, v := range sorted {
		nums[i] = v
	}
}

func main() {
	testSorting(quickSortSimple)
	fmt.Println("-=-=-=-=-=-")
	testSorting(quickSort)
}

func testSorting(f func([]int, int, int)) {
	toSort := []int{}
	f(toSort, 0, len(toSort)-1)
	fmt.Println(toSort)

	toSort = []int{1}
	f(toSort, 0, len(toSort)-1)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	f(toSort, 0, len(toSort)-1)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	f(toSort, 0, len(toSort)-1)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	f(toSort, 0, len(toSort)-1)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	f(toSort, 0, len(toSort)-1)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	f(toSort, 0, len(toSort)-1)
	fmt.Println(toSort)
}
