package main

import "fmt"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for sweepsRemaining := len(list) - 1; sweepsRemaining > 0; sweepsRemaining-- {
		swapPerformed := false
		for i := 1; i <= sweepsRemaining; i++ {
			if list[i] < list[i-1] {
				list[i-1], list[i] = list[i], list[i-1]
				swapPerformed = true
			}
		}

		if !swapPerformed {
			break
		}
	}
}

func main() {
	toSort := []int{}
	BubbleSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{1}
	BubbleSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{1, 2, 3}
	BubbleSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 2, 1}
	BubbleSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{5, 3, 5, 3, 5}
	BubbleSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{3, 5, 3, 5, 3}
	BubbleSortInt(toSort)
	fmt.Println(toSort)

	toSort = []int{10, 3, 1, 2, -4, -5, 4, 5, 7, -909, 304}
	BubbleSortInt(toSort)
	fmt.Println(toSort)
}
