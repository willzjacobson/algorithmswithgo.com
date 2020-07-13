package bf

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"math"
)

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
