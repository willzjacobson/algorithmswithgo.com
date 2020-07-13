package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	dijkstra "algo/graph-algorithms/dijkstra"
	"fmt"
)

// create vertices to be referenced below
var (
	s = adjacencylist.CreateAdjListVertex("s")
	t = adjacencylist.CreateAdjListVertex("t")
	x = adjacencylist.CreateAdjListVertex("x")
	y = adjacencylist.CreateAdjListVertex("y")
	z = adjacencylist.CreateAdjListVertex("z")
)

func buildWeightedAdjList() *adjacencylist.Weighted {
	// create adjacency list
	l := adjacencylist.CreateWeighted()
	// establish edges (important to get both directions)
	l.Adj[s] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(t, 10),
		adjacencylist.CreateAdjListEdgeWeighted(y, 5),
	}
	l.Adj[t] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(x, 1),
		adjacencylist.CreateAdjListEdgeWeighted(y, 2),
	}
	l.Adj[x] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(z, 4),
	}
	l.Adj[y] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(t, 3),
		adjacencylist.CreateAdjListEdgeWeighted(x, 9),
		adjacencylist.CreateAdjListEdgeWeighted(z, 2),
	}
	l.Adj[z] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(s, 7),
		adjacencylist.CreateAdjListEdgeWeighted(x, 6),
	}
	return l
}

// PrintShortestPath : helper to show the shortest path from the source vertex to any vertex v
func PrintShortestPath(s *adjacencylist.AdjListVertex, v *adjacencylist.AdjListVertex) {
	if v == s {
		fmt.Print("->", v.Value)
		fmt.Println()
	} else if v.P != nil {
		fmt.Print("->", v.Value)
		PrintShortestPath(s, v.P)
	} else {
		fmt.Print(" No path from source vertex s to vertex", v.Value)
		fmt.Println()
	}
}

func main() {
	l := buildWeightedAdjList()
	dijkstra.Dijkstra(l, s)
	fmt.Println("-=- Shortest paths from source vertex x to other vertices:")
	fmt.Println("Distance from s to t:", t.D)
	fmt.Println("Distance from s to x:", x.D)
	fmt.Println("Distance from s to y:", y.D)
	fmt.Println("Distance from s to z:", z.D)
	PrintShortestPath(s, t)
	PrintShortestPath(s, x)
	PrintShortestPath(s, y)
	PrintShortestPath(s, z)
}
