package helpers

import "fmt"

// AdjacencyMatrixRow : representation of a row in an AdjacencyMatrix
type AdjacencyMatrixRow []float64

// AdjacencyMatrixRowBool : used for transitive closure algorithm
type AdjacencyMatrixRowBool []bool

// AdjacencyMatrix : adjacency-matrix representation of a weighted, directed graph
type AdjacencyMatrix []AdjacencyMatrixRow

// AdjacencyMatrixBool : used for transitive closure algorithm
type AdjacencyMatrixBool []AdjacencyMatrixRowBool

// GenerateNewMatrix : helper to generate an empty nxn matrix containing value in all entries
func GenerateNewMatrix(n int, value float64) AdjacencyMatrix {
	m := AdjacencyMatrix{}
	for i := 0; i < n; i++ {
		m = append(m, AdjacencyMatrixRow{})
		for j := 0; j < n; j++ {
			m[i] = append(m[i], value)
		}
	}
	return m
}

// GenerateNewMatrixBool : helper to generate an empty nxn matrix containing false in all entries
// TODO: learn how to do generic typing, so can leverage the above method
func GenerateNewMatrixBool(n int) AdjacencyMatrixBool {
	m := AdjacencyMatrixBool{}
	for i := 0; i < n; i++ {
		m = append(m, AdjacencyMatrixRowBool{})
		for j := 0; j < n; j++ {
			m[i] = append(m[i], false)
		}
	}
	return m
}

// PrintAllPairsShortestPaths : given a precessessor graph in the form of an adjacency-matrix,
// print the shortest paths from vertex i to vertex j
func PrintAllPairsShortestPaths(π AdjacencyMatrix, i, j int) {
	if i == j {
		fmt.Println(i)
	} else if π[i][j] < 0 {
		fmt.Println("no path from", i, "to", j, "exists")
	} else {
		PrintAllPairsShortestPaths(π, i, int(π[i][j]))
		fmt.Println(j)
	}
}
