package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"fmt"
	"math"
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

// helper functions

// InitSingleSource : set up adjacency-list for Bellman-Ford algo
func InitSingleSource(l *adjacencylist.Weighted, s *adjacencylist.AdjListVertex) {
	for v := range l.Adj {
		v.D = math.Inf(1)
	}
	s.D = 0
}

// Relax : test whether the shortest path from the source to a vertex v might pass through vertex u
// If so, update the current best estimate to do that.
func Relax(u *adjacencylist.AdjListVertex, e *adjacencylist.AdjListEdgeWeighted) {
	if e.To.D > u.D+e.Weight {
		e.To.D = u.D + e.Weight
		e.To.P = u
	}
}

// BellmanFordSSSP : implementation of the Bellman-ford Single-source shortest-paths algorithm
// for generating a shortest-paths tree rooted at a single vertex
// runtime: O(V,E)
func BellmanFordSSSP(l *adjacencylist.Weighted, s *adjacencylist.AdjListVertex) bool {
	InitSingleSource(l, s)
	// Perform 'relax' procedure on each edge, V-1 times
	// This will gradually resolve the shortest-paths tree
	for i := 1; i < len(l.Adj); i++ {
		for u := range l.Adj {
			for _, e := range l.Adj[u] {
				Relax(u, e)
			}
		}
	}

	for u := range l.Adj {
		for _, e := range l.Adj[u] {
			if e.To.D > u.D+e.Weight {
				return false
			}
		}
	}

	return true
}

// PrintShortestPathsTree : helper to show the shortest path from the source vertex to any vertex v
func PrintShortestPathsTree(s, v *adjacencylist.AdjListVertex) {
	if s == v {
		fmt.Print("->", v.Value)
		fmt.Println()
		return
	}
	fmt.Print("->", v.Value)
	PrintShortestPathsTree(s, v.P)
}

func main() {
	l := buildWeightedAdjList()
	containsNoNegativeWeightCycles := BellmanFordSSSP(l, s)
	fmt.Println("The input graph contained no negative weight cycles:", containsNoNegativeWeightCycles)
	fmt.Println("-=- Shortest paths from x to other vertices:")
	PrintShortestPathsTree(s, y)
	PrintShortestPathsTree(s, x)
	PrintShortestPathsTree(s, t)
	PrintShortestPathsTree(s, z)
}
