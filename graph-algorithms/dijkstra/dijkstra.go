package dijkstra

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	minheap "algo/graph-algorithms/min-heap"
	"fmt"
	"math"
)

// helper functions

// InitSingleSource : set up adjacency-list for Bellman-Ford algo
func InitSingleSource(l *adjacencylist.Weighted, s *adjacencylist.AdjListVertex) {
	for u := range l.Adj {
		u.D = math.Inf(1)
		u.Key = u.D
	}
	s.D = 0
	s.Key = s.D
}

// Relax : test whether the shortest path from the source to a vertex v might pass through vertex u
// If so, update the current best estimate to do that.
func Relax(u *adjacencylist.AdjListVertex, e *adjacencylist.AdjListEdgeWeighted) {
	// fmt.Println("is", e.To.D, "more than", u.D, "+", e.Weight)
	if e.To.D > u.D+e.Weight {
		e.To.D = u.D + e.Weight
		e.To.P = u
	}
}

// Dijkstra : implementation of Dijkstra's algorithm for determining the shortest distance
// between any vertex and a source vertex
// Updates the parent (P) and distance (D) prop of each vertex to point to the vertex via which it joined the tree
func Dijkstra(l *adjacencylist.Weighted, s *adjacencylist.AdjListVertex) {
	InitSingleSource(l, s)
	// create a min heap and insert pointers to all the vertices
	q := minheap.CreateMinHeap()
	for u := range l.Adj {
		fmt.Println("inserting", u.Value, "with d value:", u.D)
		q.Insert(u)
	}
	for q.Size > 0 {
		u := q.ExtractMin()
		fmt.Println("extracted", u.Value, "with d value:", u.D)
		for _, e := range l.Adj[u] {
			Relax(u, e)
		}
	}
}
