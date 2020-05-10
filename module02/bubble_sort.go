package module02

import "sort"

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

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
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

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
	isLess := func(a, b Person) bool {
		if a.Age != b.Age {
			return b.Age < a.Age
		}
		if a.LastName != b.LastName {
			return a.LastName < b.LastName
		}
		return a.FirstName < b.FirstName
	}

	for sweepsRemaining := len(people) - 1; sweepsRemaining > 0; sweepsRemaining-- {
		swapPerformed := false
		for i := 1; i <= sweepsRemaining; i++ {
			if isLess(people[i], people[i-1]) {
				people[i-1], people[i] = people[i], people[i-1]
				swapPerformed = true
			}
		}

		if !swapPerformed {
			break
		}
	}
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	for sweepsRemaining := list.Len() - 1; sweepsRemaining > 0; sweepsRemaining-- {
		swapPerformed := false
		for i := 1; i <= sweepsRemaining; i++ {
			if list.Less(i, i-1) {
				list.Swap(i, i-1)
				swapPerformed = true
			}
		}

		if !swapPerformed {
			break
		}
	}
}
