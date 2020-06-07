package main

import (
	"fmt"
	"strconv"
)

/*
Matrix multiplication is associative. That is, when you have a A1 * A2 * A3, you can calculate the result via (A1 * A1) * A3, or via A1 * (A1 * A3).
The order you pick can have a huge effect on how many scalar operations you have to perform in total.
When you multiply a pxq matrix by a qxr matrix, you perform p*q*r scalar operations.

Imagine you're mutiplying 3 matrices:
	A1: 10 x 100
	A2: 100 x 5
	A3: 5 x 50

Multiplying:
	(A1 * A1) * A3
requires 10*100*5 + 10*5*50 = 7500 scalar operations, while multiplying:
	A1 * (A1 * A3)
requires 100*5*50 + 10*100*50 = 75000 scalar operations. That's 10x more!

In the matric chain multiplication problem, we're not actually multiplying matrices. Instead, we're figuring out what is the optimal way to paranthesize them for multiplication.
For a set of n matrices to be multiplied, there are 2^n ways to paranthesize them. So, this is a problem worth optimizing :)
*/

// MatrixChainOrderBottomUp : Bottom-up dynamic programming method for solving the matrix chain order problem
// There are n^2 subproblems to solve, and each one takes O(n) time. We solve each subproblem once. The full solution is O(n^3).
func MatrixChainOrderBottomUp(p []int) ([][]int, [][]int) {
	// length of the matrix chain
	n := len(p) - 1
	// build data structures:
	//  m holds # of scalar operations for each optimally paranthesized subchain
	//  s holds the optimal values of k for each optimally paranthesized sub-chain
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	s := make([][]int, n)
	for i := 0; i < n-1; i++ {
		s[i] = make([]int, n)
	}

	// To take the product of a chain of 1 matrix requires 0 scalar operations
	for i := 0; i < n; i++ {
		m[i][i] = 0
	}

	// figure out optimal values for n > 1; building the solution table bottom-up from 2
	for l := 2; l <= n; l++ { // l is length of the matrix chain in this subproblem
		// we now determine the min number of operations for each subchain of length l within the full chain
		for i := 0; i <= n-l; i++ { // since this subchain is length l, the start of the subchain can be at most n-1
			j := i + l - 1 // last matrix in the chain of length l that starts at index i
			minOperations := 0
			for k := i; k < j; k++ { // k represents a potential index at which to perenthesize this subchain
				// p[n-1] (p[n]) are the number of rows (columns) inthe nth matrix in the chain.
				// Since "i" starts at 0, so p[i] is the number of rows in the i+1'th matrix in the chain.
				// p[k+1] is the number of columns in the k+1th matrix in the chain
				// p[j+1] is the number of columns in the j+1th matrix in the chain
				q := m[i][k] + m[k+1][j] + p[i]*p[k+1]*p[j+1] // calculate the number of scalar operations if we were to perenthesize the chain between k and k+1
				if minOperations == 0 || q < minOperations {
					// so far, k is the most economical place to perenthesize the matrix chain from i-j
					minOperations = q
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	return m, s
}

// PrintOptimalPerens : Print out
func PrintOptimalPerens(s [][]int, i int, j int) string {
	if i == j {
		return "A" + strconv.Itoa(i)
	}
	start := PrintOptimalPerens(s, i, s[i][j])
	end := PrintOptimalPerens(s, s[i][j]+1, j)
	return "(" + start + end + ")"
}

// MemoizedMatrixChain : solve the matrix chain multiplication problem in a top-down recursive approach with momoization.
// As in the bottom-up approach, we have n^2 subproblems to solve, and each one takes O(n) time. We solve each subproblem once. The full solution is O(n^3).
// Without the momoization, the run time would be O(2^(n-1)), since we would solve many of the subproblems multiple times.
func MemoizedMatrixChain(p []int) int {
	// build up helper table
	n := len(p) - 1
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := range m[i] {
			m[i][j] = -1
		}
	}

	return LookupChain(p, 0, n-1, m)
}

// LookupChain : helper method for MatrixChainOrderTopDown
func LookupChain(p []int, i, j int, m [][]int) int {
	if m[i][j] != -1 {
		return m[i][j]
	} else if i == j {
		return 0
	} else {
		for k := i; k < j; k++ {
			q := LookupChain(p, i, k, m) + LookupChain(p, k+1, j, m) + p[i]*p[k+1]*p[j+1]
			if m[i][j] == -1 || q < m[i][j] {
				m[i][j] = q
			}
		}
		return m[i][j]
	}
}

func main() {
	p := []int{30, 35, 15, 5, 10, 20, 25}
	m, s := MatrixChainOrderBottomUp(p)
	fmt.Println("m:", m)
	fmt.Println("s:", s)
	perens := PrintOptimalPerens(s, 0, len(p)-2)
	fmt.Println("perens:", perens)

	// Memoized top down approach
	n := MemoizedMatrixChain(p)
	fmt.Println("Using top down memoized approach:", n)
}
