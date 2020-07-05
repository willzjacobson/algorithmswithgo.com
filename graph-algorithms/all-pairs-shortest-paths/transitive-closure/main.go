package main

import (
	"algo/graph-algorithms/all-pairs-shortest-paths/helpers"
	"fmt"
)

func generateStartingAdjacencyMatrix1() helpers.AdjacencyMatrixBool {
	return helpers.AdjacencyMatrixBool{
		{true, true, true, false, true},
		{false, true, false, true, true},
		{false, true, true, false, false},
		{true, false, true, true, false},
		{false, false, false, true, true},
	}
}

func generateStartingAdjacencyMatrix2() helpers.AdjacencyMatrixBool {
	return helpers.AdjacencyMatrixBool{
		{true, false, false, false},
		{false, false, true, true},
		{false, true, false, false},
		{true, false, true, false},
	}
}

// TransitiveClosure : generate a matrix t that reports whether a path exists between each pair of vertices
// this could easily be modified to accept an adjacency-list rather than an adjacency-matrix
func TransitiveClosure(g helpers.AdjacencyMatrixBool) helpers.AdjacencyMatrixBool {
	n := len(g)

	// build initial t matrix
	// given that this method is being passed an AdjacencyMatrixBool, this part is not necessary.
	// but something similar would be necessary if we were being passed an adjacency-list or a regular adjacency-matrix
	t := helpers.GenerateNewMatrixBool(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || g[i][j] {
				t[i][j] = true
			} else {
				t[i][j] = false
			}
		}
	}

	// build upon t by adding 1 new vertex k to the set of allowed intermediate vertices in a path from i->j
	for k := 0; k < n; k++ {
		tk := helpers.GenerateNewMatrixBool(n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if t[i][j] || (t[i][k] && t[k][j]) {
					tk[i][j] = true
				}
			}
		}
		t = tk
	}

	return t
}

func main() {
	// from figure 25.1
	g1 := generateStartingAdjacencyMatrix1()
	tranClosure1 := TransitiveClosure(g1)
	fmt.Println("-=- Transitive Closure Result 1:")
	fmt.Println(tranClosure1)

	// from figure 25.5
	g2 := generateStartingAdjacencyMatrix2()
	tranClosure2 := TransitiveClosure(g2)
	fmt.Println("-=- Transitive Closure Result 2:")
	fmt.Println(tranClosure2)
}
