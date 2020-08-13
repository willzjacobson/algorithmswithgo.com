package main

import "fmt"

/*
	In the longest-common-subsequence problem, we are given two sequences, and wish to find a maximum-length common subsequence.
	A shorter string Y is a subsequence of a longer string X if you can obtain Y simply by removing characters from X, without changing the order of the remaining characters.

	If we want to find the LCS (longest common subsequence) between 2 strings X and Y, a brute force approach would enumarate through all subsequences of X, and see if they are also contained in Y.
	If X has length m, X contains 2^m subsequences. And then we have to look for them in Y. Ouch.

	The LCS problem has an optimal subscructure for dynamic programming. That is because an LCS of 2 sequences contains within it an LCS of prefixes of the 2 sequences.
	If X has length m and Y has length n:
	 - if Xm=Yn, we then solve the same problem for finding the LCS for X1-Xm-1 and Y1-Yn-1 (which is a subproblem)
	 - if Xm!=Yn, we then solve the 2 subproblems: LCS of X1-Xm and Y1-Yn-1, and LCS of X1-Xm-1 and Y1-Yn
	We see that there are m*n subproblems. We could create a recursive approach that solves some of the subproblems several times over. Instead, we use dynamic programming to compute the solution bottom-up.
*/

const (
	left   = "left"
	leftUp = "leftUp"
	up     = "up"
)

// LongestSubstring : solve the longest substring problem using a bottom-up approach with dynamic programming
func LongestSubstring(X, Y string) ([][]int, [][]string) {
	m := len(X)
	n := len(Y)
	// initialize table of all 0's
	c := make([][]int, m+1)    // a table of LCS length for each subproblem
	b := make([][]string, m+1) // a table used to reconstruct the LCS
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			c[i] = append(c[i], 0)
			b[i] = append(b[i], "0")
		}
	}

	for i, x := range X {
		for j, y := range Y {
			// All the '+1' stuff is to account for indexing starting at 0, while maintaing a 0th row and column in the table before the string starts
			if x == y {
				c[i+1][j+1] = c[i][j] + 1
				b[i+1][j+1] = leftUp
			} else if c[i][j+1] >= c[i+1][j] {
				c[i+1][j+1] = c[i][j+1]
				b[i+1][j+1] = up // cutting off last char of X
			} else {
				c[i+1][j+1] = c[i+1][j]
				b[i+1][j+1] = left // cutting off last char of Y
			}
		}
	}

	return c, b
}

// PrintLCS : helper to reconstruct the LCS of A and Y using table b
// Takes O(m+n) time, where m is length of X and
func PrintLCS(X, Y string, b [][]string) string {
	lcs := ""
	i := len(X)
	j := len(Y)
	for i > 0 && j > 0 {
		if b[i][j] == leftUp {
			lcs = string(X[i-1]) + lcs
			i--
			j--
		} else if b[i][j] == left {
			j--
		} else {
			i--
		}
	}
	return lcs
}

func main() {
	x := "ABCBDAB"
	y := "BDCABA"
	c, b := LongestSubstring(x, y)
	fmt.Println("c:", c)
	fmt.Println("b:", b)
	lcs := PrintLCS(x, y, b)
	fmt.Println("LCS:", lcs)
}
