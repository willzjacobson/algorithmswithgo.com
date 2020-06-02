/*
A company buys rods of length x (integral # of inches), cuts them into n pieces, and sells the pieces. The process of cutting is free.
A rod of i inches costs a different amount. The company needs to know the most profitable way to cut the rods.

We can cut a rod of length n 2^(n-1) different ways (since at each inch mark, we can either cut or not cut, thus generating another 2 potential outcomes).
*/

package main

import (
	"fmt"
)

// **** Divide And Conquer Method

/*
CutRodRecursive : Recursive solution using divide-and-conquer methodology.
Takes an input p[0:n] pf prices and an integer n (length of the rod).
Returns the maximum possible revenue for a rod of length n.
This is what we'd call a "top down" approach, since we start with the full rod (the full problem) and break it down into parts
*/
func CutRodRecursive(p []int, n int) int {
	var q int = 0 // keeping track of the max

	// a rod of length 0 has no price
	if n == 0 {
		return q
	}

	for i := 1; i <= n; i++ {
		maxResultFromFollowingThisSubtree := p[i] + CutRodRecursive(p, n-i)

		if maxResultFromFollowingThisSubtree > q {
			q = maxResultFromFollowingThisSubtree
		}
	}

	return q
}

// **** Dynamic Programming Methods

// CutRodTopDown : "top-down with memoization"
// Store the resulting max for a rod of length i whenever it's encountered
// O(n) due to having ~n recursive calls (that return values from the cache) in each round of the loop
func CutRodTopDown(p []int, n int) int {
	cache := map[int]int{}
	for i := 0; i <= n; i++ {
		cache[i] = -1
	}
	return CutRodTopDownAux(p, n, cache)
}

func CutRodTopDownAux(p []int, n int, c map[int]int) int {
	if n == 0 {
		return 0
	}

	if c[n] != -1 {
		return c[n]
	}

	q := 0
	for i := 1; i <= n; i++ {
		maxResultFromFollowingThisSubtree := p[i] + CutRodTopDownAux(p, n-i, c)
		if maxResultFromFollowingThisSubtree > q {
			q = maxResultFromFollowingThisSubtree
		}
	}
	c[n] = q
	return q
}

// CutRodBottomUp : Top-down with memoization
// This approach requires some notion of the "size" of a subproblem, sich that solving any particular subproblem depends only on solving "smaller" subproblems.
// This method generally has better constant factors, since it has less overhead for procedure calls
// O(n) due to nexted loop
func CutRodBottomUp(p []int, n int) int {
	c := []int{0} // cache

	for i := 1; i <= n; i++ {
		maxResultForRodOfLengthI := 0
		for j := 1; j <= i; j++ {
			maxResultFromFollowingThisSubtree := p[j] + c[i-j]
			if maxResultFromFollowingThisSubtree > maxResultForRodOfLengthI {
				maxResultForRodOfLengthI = maxResultFromFollowingThisSubtree
			}
		}
		c = append(c, maxResultForRodOfLengthI)
	}
	return c[n]
}

// CutRodBottomUpRecordResult : same as above, but recording out the length of the first rod cut
func CutRodBottomUpRecordResult(p []int, n int) (int, int) {
	c := []int{0}         // cache
	s := make([]int, n+1) // list of first cut length for each starting length

	for i := 1; i <= n; i++ {
		maxResultForRodOfLengthI := 0
		for j := 1; j <= i; j++ {
			maxResultFromFollowingThisSubtree := p[j] + c[i-j]
			if maxResultFromFollowingThisSubtree > maxResultForRodOfLengthI {
				maxResultForRodOfLengthI = maxResultFromFollowingThisSubtree
				s[i] = j
			}
		}
		c = append(c, maxResultForRodOfLengthI)
	}
	return c[n], s[n]
}

// CutRodBottomUpPrintResult : Prints results from CutRodBottomUpPrintResult
func CutRodBottomUpPrintResult(p []int, n int) {
	for i := 1; i <= n; i++ {
		v, s := CutRodBottomUpRecordResult(p, i)
		fmt.Println("Length:", i, " | value:", v, "first Selection:", s)
	}
}

// ** Testing
func main() {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	fmt.Println("-=- Generated Using Recursive Solution:")
	fmt.Println("0:", CutRodRecursive(p, 0))
	fmt.Println("1:", CutRodRecursive(p, 1))
	fmt.Println("2:", CutRodRecursive(p, 2))
	fmt.Println("3:", CutRodRecursive(p, 3))
	fmt.Println("4:", CutRodRecursive(p, 4))
	fmt.Println("5:", CutRodRecursive(p, 5))
	fmt.Println("6:", CutRodRecursive(p, 6))
	fmt.Println("7:", CutRodRecursive(p, 7))
	fmt.Println("8:", CutRodRecursive(p, 8))
	fmt.Println("9:", CutRodRecursive(p, 9))
	fmt.Println("10:", CutRodRecursive(p, 10))

	fmt.Println("-=- Generated Using Dynamic Programming Top-Down Solution:")
	fmt.Println("0:", CutRodTopDown(p, 0))
	fmt.Println("1:", CutRodTopDown(p, 1))
	fmt.Println("2:", CutRodTopDown(p, 2))
	fmt.Println("3:", CutRodTopDown(p, 3))
	fmt.Println("4:", CutRodTopDown(p, 4))
	fmt.Println("5:", CutRodTopDown(p, 5))
	fmt.Println("6:", CutRodTopDown(p, 6))
	fmt.Println("7:", CutRodTopDown(p, 7))
	fmt.Println("8:", CutRodTopDown(p, 8))
	fmt.Println("9:", CutRodTopDown(p, 9))
	fmt.Println("10:", CutRodTopDown(p, 10))

	fmt.Println("-=- Generated Using Dynamic Programming Bottom-Up Solution:")
	fmt.Println("0:", CutRodBottomUp(p, 0))
	fmt.Println("1:", CutRodBottomUp(p, 1))
	fmt.Println("2:", CutRodBottomUp(p, 2))
	fmt.Println("3:", CutRodBottomUp(p, 3))
	fmt.Println("4:", CutRodBottomUp(p, 4))
	fmt.Println("5:", CutRodBottomUp(p, 5))
	fmt.Println("6:", CutRodBottomUp(p, 6))
	fmt.Println("7:", CutRodBottomUp(p, 7))
	fmt.Println("8:", CutRodBottomUp(p, 8))
	fmt.Println("9:", CutRodBottomUp(p, 9))
	fmt.Println("10:", CutRodBottomUp(p, 10))

	fmt.Println("-=- Dynamic Programming Bottom-Up Solution, with first rod length:")
	CutRodBottomUpPrintResult(p, 10)
}
