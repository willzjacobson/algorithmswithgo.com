package main

import "fmt"

func selectionSort(nums []int) {
	for i := range nums {
		min := nums[i]
		minInd := i
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				minInd = j
			}
		}
		nums[i], nums[minInd] = nums[minInd], nums[i]
	}
}

func main() {
	toSort := []int{}
	selectionSort(toSort)
	fmt.Println(toSort)

	toSort = []int{1}
	selectionSort(toSort)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	selectionSort(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	selectionSort(toSort)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	selectionSort(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	selectionSort(toSort)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	selectionSort(toSort)
	fmt.Println(toSort)
}
