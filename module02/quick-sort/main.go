package main

import "fmt"

func quickSort(nums []int) {
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
	quickSort(lower)
	quickSort(higher)

	// compose fully sorted array
	sorted := append(lower, append([]int{mid}, higher...)...)
	// overlay sorted order onto original array
	for i, v := range sorted {
		nums[i] = v
	}
}

func main() {
	toSort := []int{}
	quickSort(toSort)
	fmt.Println(toSort)

	toSort = []int{1}
	quickSort(toSort)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	quickSort(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	quickSort(toSort)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	quickSort(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	quickSort(toSort)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	quickSort(toSort)
	fmt.Println(toSort)
}
