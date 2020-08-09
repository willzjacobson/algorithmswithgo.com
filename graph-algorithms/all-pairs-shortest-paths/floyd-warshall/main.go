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

// constructPredecessorSubgraph : mutates π in place
func constructPredecessorSubgraph(dkMin1 helpers.AdjacencyMatrix, dk helpers.AdjacencyMatrix, π helpers.AdjacencyMatrix, k int) helpers.AdjacencyMatrix {
	n := len(dkMin1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dk[i][j] != dkMin1[i][j] {
				// the value of this entry of d(k) is different from that of d(k-1); the shortest path has changed
				if k < 0 {
					// just initializing π; if i != j, the predecessor of j must be i
					if i != j {
						π[i][j] = float64(i)
					}
				} else {
					// new predecessor of j in the path i->j is the same as the predecessor of j in the path k->j
					π[i][j] = π[k][j]
				}
			}
		}
	}
	return π
}

// FloydWarshall : generate all-pairs shortest-paths matrix in O(V^3) time
func FloydWarshall(w helpers.AdjacencyMatrix) (helpers.AdjacencyMatrix, helpers.AdjacencyMatrix) {
	n := len(w) // number of vertices in G
	// initialize shortest paths matrix
	dk := w
	// initialize predecessor matrix
	πk := constructPredecessorSubgraph(
		helpers.GenerateNewMatrix(n, math.Inf(1)),
		w,
		helpers.GenerateNewMatrix(n, -1), // initialize π as all nil entries
		-1,                               // k=-1 indicates we are just initializing π
	)

	// outer loop adds a vertex to the set of allowed intermediate vertices for our growing shortest-paths matrix
	for k := 0; k < n; k++ {
		d := helpers.GenerateNewMatrix(n, math.Inf(1))
		// inner loops compute the entry of the new shortest-paths matrix d(k)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// use the value from k(k-1) unless adding vertex k as an intermediate vertex yields a lower-weight path
				d[i][j] = math.Min(dk[i][j], dk[i][k]+dk[k][j])
			}
		}

		// compute new π(k) matrix using latest shortest-paths matrix d(k)
		πk = constructPredecessorSubgraph(dk, d, πk, k)
		dk = d // update latest shortest-paths matrix
	}

	// return both the shortest-paths matrix and the predeccessor graph
	return dk, πk
}

func main() {
	w := generateStartingAdjacencyMatrix()

	shortestPaths, predecessorGraph := FloydWarshall(w)
	fmt.Println("-=- Floyd Warshall APSP result:")
	fmt.Println(shortestPaths)
	fmt.Println("-=- Predecessor Graph:")
	fmt.Println(predecessorGraph)
}
