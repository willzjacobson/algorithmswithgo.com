package main

import "fmt"

// CountingSort will sort a list of integers using the counting sort algorithm.
//
// Big O: O(k+n), where n is the size of the list containing integers in the range 0-k
func CountingSort(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	out := make([]int, len(nums))
	store := make([]int, k+1)

	// each elem in store now contains the number of elements equal to v
	for _, v := range nums {
		store[v]++
	}
	// each elem in store now contains the number of elements <= v
	for i := 1; i < len(store); i++ {
		store[i] += store[i-1]
	}
	// place each element in nums in the correct place in the output array
	for _, v := range nums {
		out[store[v]-1] = v
		store[v]--
	}

	for i, v := range out {
		nums[i] = v
	}
}

func main() {
	toSort := []int{}
	CountingSort(toSort, 0)
	fmt.Println(toSort)

	toSort = []int{1}
	CountingSort(toSort, 1)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	CountingSort(toSort, 3)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	CountingSort(toSort, 3)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	CountingSort(toSort, 5)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	CountingSort(toSort, 5)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, 4, 5, 7, 5, 2}
	CountingSort(toSort, 10)
	fmt.Println(toSort)
}
