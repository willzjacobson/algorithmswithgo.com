package dfs

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
)

const (
	white = "white"
	gray  = "gray"
	black = "black"
)

// universal time keeper for DFS
var time = 0

// DFS : DFS implementation on a graph modelled as an adjacency-list
// The resulting depth-first forest will actually come out slightly differently
// depending on the order of the keys adjacency-list iterates through which the loop iterates
func DFS(g *adjacencylist.AdjacencyList) {
	// set up adjacency-list
	for v := range g.Adj {
		v.Color = white
	}

	// loop through all vertices, and process the ones that are still white
	// (not all vertices may be processed by Visit at this callsite,
	// since some will be processed by recursive calls of Visit resulting from the initial call(s))
	for k := range g.Adj {
		if k.Color == white {
			Visit(g, k)
		}
	}
}

// Visit : Helper for DFS
func Visit(g *adjacencylist.AdjacencyList, u *adjacencylist.AdjListVertex) {
	// vertex u is being discovered now: mark it gray and add "start" timestamp
	time++
	u.Start = time
	u.Color = gray

	// loop through the vertices adjacent to u and process any that are still white
	for _, v := range g.Adj[u] {
		if v.Color == white {
			v.P = u
			Visit(g, v)
		}
	}

	// vertex u is now finished: mark it black and add "end" timestamp
	time++
	u.End = time
	u.Color = black
}
