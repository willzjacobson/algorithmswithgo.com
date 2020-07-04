package main

import (
	"fmt"
	"math"
)

type adjacencyMatrixRow []float64
type adjacencyMatrix []adjacencyMatrixRow

func generateStartingAdjacencyMatrix() adjacencyMatrix {
	return adjacencyMatrix{
		{0, 3, 8, math.Inf(1), -4},
		{math.Inf(1), 0, math.Inf(1), 1, 7},
		{math.Inf(1), 4, 0, math.Inf(1), math.Inf(1)},
		{2, math.Inf(1), -5, 0, math.Inf(1)},
		{math.Inf(1), math.Inf(1), math.Inf(1), 6, 0},
	}
}

// helper to generate an empty nxn matrix containing +∞ in all entries
func generateNewMatrix(n int) adjacencyMatrix {
	m := adjacencyMatrix{}
	for i := 0; i < n; i++ {
		m = append(m, adjacencyMatrixRow{})
		for j := 0; j < n; j++ {
			m[i] = append(m[i], math.Inf(1))
		}
	}
	return m
}

func extendShortestPaths(l adjacencyMatrix, w adjacencyMatrix) adjacencyMatrix {
	// generate a new matrix that will contain the extended shortest paths
	n := len(l)
	lNext := generateNewMatrix(n)
	// determine the value of each entry of the new shortest paths matrix (requies 3 nested loops, like matrix multiplication)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				// The value is the min of the previous shortest path weight (if one exists)
				// and the path weight generated using all possible predecessors of j
				lNext[i][j] = math.Min(lNext[i][j], l[i][k]+w[k][j])
			}
		}
	}
	fmt.Println("-=- Next shortest paths matrix:")
	// fmt.Println(lNext)
	return lNext
}

// O(V^4) implementation
func slowAllPairsShortestPaths(w adjacencyMatrix) adjacencyMatrix {
	n := len(w)
	l := w
	// loop |E|-2 times (since the shortest paths will be m=|E|-1 at longest,
	// and w already gives us the all-pairs shortest-paths matrix for m=1)
	for i := 2; i < n; i++ {
		l = extendShortestPaths(l, w)
	}
	return l
}

// O(V^4) implementation
func fasterAllPairsShortestPaths(w adjacencyMatrix) adjacencyMatrix {
	n := len(w)
	l := w
	m := 1
	// This time, loop until m > n-1, doubling m each time through the loop
	// Total of lg n-1 rounds, which is better than n-2
	for m < n-1 {
		l = extendShortestPaths(l, l)
		m *= 2
	}
	return l
}

func main() {
	w := generateStartingAdjacencyMatrix()

	APSPSlow := slowAllPairsShortestPaths(w)
	fmt.Println("-=- result from O(V^4) algorithm:")
	fmt.Println(APSPSlow)
	fmt.Println("-=-=-=-")
	APSPFaster := fasterAllPairsShortestPaths(w)
	fmt.Println("-=- result from O(lgV V^3) algorithm:")
	fmt.Println(APSPFaster)
}
