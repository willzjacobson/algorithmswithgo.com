package main

import (
	"algo/graph-algorithms/all-pairs-shortest-paths/helpers"
	"fmt"
	"math"
)

func generateStartingAdjacencyMatrix() helpers.AdjacencyMatrix {
	return helpers.AdjacencyMatrix{
		{0, 3, 8, math.Inf(1), -4},
		{math.Inf(1), 0, math.Inf(1), 1, 7},
		{math.Inf(1), 4, 0, math.Inf(1), math.Inf(1)},
		{2, math.Inf(1), -5, 0, math.Inf(1)},
		{math.Inf(1), math.Inf(1), math.Inf(1), 6, 0},
	}
}

func extendShortestPaths(m helpers.AdjacencyMatrix, w helpers.AdjacencyMatrix) helpers.AdjacencyMatrix {
	// generate a new matrix that will contain the extended shortest paths
	n := len(m)
	mNext := helpers.GenerateNewMatrix(n, math.Inf(1))
	fmt.Println("-=- GO", m[4])
	// determine the value of each entry of the new shortest paths matrix (requies 3 nested loops, like matrix multiplication)
	for i := 4; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				// The value is the min of the previous shortest path weight (if one exists)
				// and the path weight generated using all possible predecessors of j
				mNext[i][j] = math.Min(m[i][j], mNext[i][k]+w[k][j])
				if j == 0 {
					fmt.Println("HELLO from ", k, m[i][j], mNext[i][k]+w[k][j], mNext[i][j])
				}
				// fmt.Println(i, j, k, "| was:", m[i][j], "now:", mNext[i][j])
			}
		}
	}
	fmt.Println("should be 8:", mNext[4][0])
	return mNext
}

// O(V^4) implementation
func slowAllPairsShortestPaths(w helpers.AdjacencyMatrix) helpers.AdjacencyMatrix {
	n := len(w)
	m := w
	// loop |E|-2 times (since the shortest paths will be m=|E|-1 at longest,
	// and w already gives us the all-pairs shortest-paths matrix for m=1)
	for i := 2; i < n; i++ {
		m = extendShortestPaths(m, w)
		fmt.Println("AFTER ROUND", n-2, ": ", m[4])
		break
	}
	return m
}

// O(lgV V^3) implementation
func fasterAllPairsShortestPaths(w helpers.AdjacencyMatrix) helpers.AdjacencyMatrix {
	n := len(w)
	m := w
	kk := 1
	// This time, loop until m > n-1, doubling m each time through the loop
	// Total of lg n-1 rounds, which is better than n-2
	for kk < n-1 {
		m = extendShortestPaths(m, m)
		kk *= 2
	}
	return m
}

func main() {
	w := generateStartingAdjacencyMatrix()

	APSPSlow := slowAllPairsShortestPaths(w)
	fmt.Println("-=- result from O(V^4) algorithm:")
	fmt.Println(APSPSlow)
	fmt.Println("-=-=-=-")
	// APSPFaster := fasterAllPairsShortestPaths(w)
	// fmt.Println("-=- result from O(lgV V^3) algorithm:")
	// fmt.Println(APSPFaster)
}
