package main

import "fmt"

/*
	We have a list of activities that have a start time and end time. They are in monotomically increasing order of finish time. All the activities use a common resource (for example, a conference room). 2 activities ai and aj are compatible if either startj >= endi or starti >= endj. We want to work out the ideal selection of activities such that we fit the highest number of them in.

	This problem has optimal substructure for dynamic programming. However, we can actually choose an activity to add to our optimal solution without solving all the subproblems. For this problem, we only need to consider 1 choice: the greedy choice. At each juncture, we can simply select the activity that has the earliest finish time. At the start, we select a1, since they are sorted by finish time. Our only subproblem, then, is to find that activity within the subset of activities that start after the previously selected one finishes.

	A greedy algo like this does not need to work bottom up, filling in a table as it goes. It can work top down, adding to an optimal solution, then doing the same for subproblems of decreasing size.
*/

// RecursiveGreedyActivitySelector : recursive, top-down approach
// O(n) time
func RecursiveGreedyActivitySelector(s, f []int, lastActivityIndexChosen, numActivities int) []int {
	// if this is the first round, select the first activity since we know it has the earliest finish time
	if lastActivityIndexChosen == -1 {
		return append([]int{0}, RecursiveGreedyActivitySelector(s, f, 0, numActivities)...)
	}

	// find the next activity with the earliest finish time that starts after the previously selected one ends
	nextActivityIndex := lastActivityIndexChosen + 1
	for nextActivityIndex < numActivities && s[nextActivityIndex] < f[lastActivityIndexChosen] {
		nextActivityIndex++
	}

	// recursvely build the solution
	if nextActivityIndex < numActivities {
		return append([]int{nextActivityIndex}, RecursiveGreedyActivitySelector(s, f, nextActivityIndex, numActivities)...)
	}

	// no more compatible activities
	return []int{}
}

// GreedyActivitySelector : iterative (preferable), top-down approach
// O(n) time
func GreedyActivitySelector(s, f []int) []int {
	n := len(s)
	solution := []int{0}

	lastChosen := 0
	for i := lastChosen + 1; i < n; i++ {
		if s[i] >= f[lastChosen] {
			solution = append(solution, i)
			lastChosen = i
		}
	}

	return solution
}

func main() {
	startTimes := []int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	endTimes := []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}
	// pass in -1 as lastActivityIndexChosen so the problem works on the first round
	maxActivitySet1 := RecursiveGreedyActivitySelector(startTimes, endTimes, -1, len(startTimes))
	fmt.Println("maxActivitySet from recursive solution:", maxActivitySet1)

	maxActivitySet2 := GreedyActivitySelector(startTimes, endTimes)
	fmt.Println("maxActivitySet from recursive solution:", maxActivitySet2)
}
