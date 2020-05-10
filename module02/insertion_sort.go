package module02

import (
	"sort"
)

// InsertionSortIntBrute will sort a list of integers using the insertion sort
// algorithm, using brute force rather than binary search
// to find the index at which to insert each element in the sorted array
//
// Big O (without binary search): O(N^2), where N is the size of the list.
//
// See below for binary search implementation
func InsertionSortIntBrute(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insertInt(sorted, item)
	}

	for i, v := range sorted {
		list[i] = v
	}
}

func insertInt(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			return append(sorted[:i], append([]int{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item)
}

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

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
	var sorted []string
	for _, item := range list {
		sorted = insertString(sorted, item)
	}

	for i, v := range sorted {
		list[i] = v
	}
}

func insertString(sorted []string, item string) []string {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			return append(sorted[:i], append([]string{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item)
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
	var sorted []Person
	for _, item := range people {
		sorted = insertPerson(sorted, item)
	}

	for i, v := range sorted {
		people[i] = v
	}
}

func insertPerson(sorted []Person, person Person) []Person {
	for i, sortedPerson := range sorted {
		if isLessPerson(person, sortedPerson) {
			return append(sorted[:i], append([]Person{person}, sorted[i:]...)...)
		}
	}
	return append(sorted, person)
}

func isLessPerson(a, b Person) bool {
	if a.Age < b.Age {
		return true
	}
	if a.LastName < b.LastName {
		return true
	}
	return a.FirstName < b.FirstName
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		// list will always be sorted up to j
		// i is the next element to sort
		for j := 0; j < i; j++ {
			// figuring out where i will go
			if list.Less(i, j) {
				// swap i with i-1 intil i gets swapped with j, and j winds up in j+1
				for k := i; k > j; k-- {
					list.Swap(k, k-1)
				}
			}
		}
	}
}
