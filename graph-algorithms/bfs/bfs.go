package bfs

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"fmt"
)

const (
	white = "white"
	gray  = "gray"
	black = "black"
)

// BFS : breadth-first search on an adjacency list implementation of a graph
func BFS(g *adjacencylist.AdjacencyList, s *adjacencylist.AdjListVertex) {
	// set up adjacency-list for this algorithm
	for _, v := range g.Adj {
		for _, e := range v {
			e.Color = white // all vertices start white
			e.D = -1        // distance from source initialized to -1 as a null value
		}
	}
	// set properties on s and add it to the queue
	s.Color = gray // nodes are gray if in the queue
	s.D = 0        // distance from s to s is 0
	Q := []*adjacencylist.AdjListVertex{s}

	// perform BFS
	for len(Q) > 0 {
		// process first vertex in the queue
		v := Q[0]
		Q = Q[1:]
		// add properties to nodes connected to v, and add them to the queue for processing on the next round
		for _, e := range g.Adj[v] {
			if e.Color == white {
				e.P = v
				e.D = v.D + 1
				e.Color = gray
				Q = append(Q, e)
			}
		}
		v.Color = black
	}
}

// PrintTree : Print a shortest path from the source vertex of a graph to any node
// First param g must be the result of running BFS on a graph modelled by an adjacency list
func PrintTree(g *adjacencylist.AdjacencyList, s, v *adjacencylist.AdjListVertex) {
	if v == s {
		fmt.Println(s)
	} else if v.P != nil {
		PrintTree(g, s, v.P)
		fmt.Println(v)
	} else {
		fmt.Println("A path between these two vertices does not exist :(")
	}
}
