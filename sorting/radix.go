package main

import (
	"fmt"
)

// COME BACK TO THIS
func RadixSort(nums []string, d int) {
	for i := d - 1; i >= 0; i-- {

	}
}

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
	toSort := []string{"178", "399", "112", "202", "434", "583", "723", "533", "201"}
	RadixSort(toSort, 3)
	fmt.Println(toSort)
}
