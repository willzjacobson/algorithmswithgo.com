package main

import "fmt"

func mergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	midInd := len(nums) / 2
	left := append([]int(nil), nums[:midInd]...)
	right := append([]int(nil), nums[midInd:]...)

	mergeSort(left)
	mergeSort(right)

	merge(nums, left, right)
}

func merge(nums, left, right []int) {
	lenLeft := len(left)
	leftInd := 0
	lenRight := len(right)
	rightInd := 0

	totalLen := lenLeft + lenRight
	currentInd := 0

	// Do the merging
	for currentInd < totalLen {
		if leftInd == lenLeft {
			nums[currentInd] = right[rightInd]
			rightInd++
		} else if rightInd == lenRight {
			nums[currentInd] = left[leftInd]
			leftInd++
		} else if left[leftInd] < right[rightInd] {
			nums[currentInd] = left[leftInd]
			leftInd++
		} else {
			nums[currentInd] = right[rightInd]
			rightInd++
		}
		// one more element has been merged
		currentInd++
	}
}

func main() {
	toSort := []int{}
	mergeSort(toSort)
	fmt.Println(toSort)

	toSort = []int{1}
	mergeSort(toSort)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	mergeSort(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	mergeSort(toSort)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	mergeSort(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	mergeSort(toSort)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	mergeSort(toSort)
	fmt.Println(toSort)
}
