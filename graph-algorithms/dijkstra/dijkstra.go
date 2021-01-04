package dijkstra

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	minheap "algo/graph-algorithms/min-heap"
	"math"
)

// helper functions

// InitSingleSource : set up adjacency-list
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
func Relax(u *adjacencylist.AdjListVertex, e *adjacencylist.AdjListEdgeWeighted, q *minheap.MinHeap) {
	if e.To.D > u.D+e.Weight {
		e.To.D = u.D + e.Weight
		e.To.P = u
		q.DecreaseKey(e.To.Index, e.To.D)
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
		q.Insert(u)
	}
	for q.Size > 0 { // O(V) cost, since each vertex is in the queue
		u := q.ExtractMin()          // O(lg V) cost
		for _, e := range l.Adj[u] { // O(E) cost if you include the outer for-loop since looping through all edges for each vertex
			Relax(u, e, q) // O(lg V) cost due to q.decreaseKey call (would be O(1) with Fibonacci-heap)
		}
	}
}
