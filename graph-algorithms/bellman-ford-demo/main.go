package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	bf "algo/graph-algorithms/bellman-ford"
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
	// establish edges
	l.Adj[s] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(t, 6),
		adjacencylist.CreateAdjListEdgeWeighted(y, 7),
	}
	l.Adj[t] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(x, 5),
		adjacencylist.CreateAdjListEdgeWeighted(y, 8),
		adjacencylist.CreateAdjListEdgeWeighted(z, -4),
	}
	l.Adj[x] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(t, -2),
	}
	l.Adj[y] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(x, -3),
		adjacencylist.CreateAdjListEdgeWeighted(z, 9),
	}
	l.Adj[z] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(s, 2),
		adjacencylist.CreateAdjListEdgeWeighted(x, 7),
	}
	return l
}

// PrintShortestPath : helper to show the shortest path from the source vertex to any vertex v
func PrintShortestPath(s, v *adjacencylist.AdjListVertex) {
	if s == v {
		fmt.Print("->", v.Value)
		fmt.Println()
		return
	}
	fmt.Print("->", v.Value)
	PrintShortestPath(s, v.P)
}

func main() {
	l := buildWeightedAdjList()
	containsNoNegativeWeightCycles := bf.BellmanFordSSSP(l, s)
	fmt.Println("The input graph contained no negative weight cycles:", containsNoNegativeWeightCycles)
	fmt.Println("-=- Shortest paths from x to other vertices:")
	PrintShortestPath(s, y)
	PrintShortestPath(s, x)
	PrintShortestPath(s, t)
	PrintShortestPath(s, z)
}
