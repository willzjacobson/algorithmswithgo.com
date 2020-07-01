package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"fmt"
	"math"
)

// create vertices to be referenced below
var (
	r = adjacencylist.CreateAdjListVertex("r")
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
	l.Adj[r] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(s, 5),
		adjacencylist.CreateAdjListEdgeWeighted(t, 3),
	}
	l.Adj[s] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(t, 2),
		adjacencylist.CreateAdjListEdgeWeighted(x, 6),
	}
	l.Adj[t] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(x, 7),
		adjacencylist.CreateAdjListEdgeWeighted(z, 2),
		adjacencylist.CreateAdjListEdgeWeighted(y, 4),
	}
	l.Adj[x] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(y, -1),
		adjacencylist.CreateAdjListEdgeWeighted(z, 1),
	}
	l.Adj[y] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(z, -2),
	}
	l.Adj[z] = []*adjacencylist.AdjListEdgeWeighted{}
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

// DagShortestPaths : generate a shortest-paths tree for a graph repped by an adjacency-list, given any source node s
// runtime: O(V+E)
func DagShortestPaths(l *adjacencylist.Weighted, s *adjacencylist.AdjListVertex) {
	// Init
	InitSingleSource(l, s)

	// Topologically sort vertcies, and print them to verify
	TopologicalSort(l)
	for _, v := range sorted {
		fmt.Print(v.Value, ",", v.End, " -> ")
	}
	fmt.Println()

	for _, v := range sorted {
		for _, e := range l.Adj[v] {
			Relax(v, e)
		}
	}
}

// TOPOLOGICAL SORT IMPLEMENTATION
const (
	white = "white"
	gray  = "gray"
	black = "black"
)

// Universal time keeper for DFS
var time = 0

// Sorted list of vertices
var sorted = []*adjacencylist.AdjListVertex{}

// TopologicalSort : perform topological sort
func TopologicalSort(g *adjacencylist.Weighted) {
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
func Visit(g *adjacencylist.Weighted, u *adjacencylist.AdjListVertex) {
	// vertex u is being discovered now: mark it gray and add "start" timestamp
	time++
	u.Start = time
	u.Color = gray

	// loop through the vertices adjacent to u and process any that are still white
	for _, v := range g.Adj[u] {
		if v.To.Color == white {
			v.To.P = u
			v.To.Weight = v.Weight
			Visit(g, v.To)
		}
	}

	// vertex u is now finished: mark it black and add "end" timestamp
	time++
	u.End = time
	u.Color = black
	sorted = append([]*adjacencylist.AdjListVertex{u}, sorted...)
}

// PrintShortestPathsTree : helper to show the shortest path from the source vertex to any vertex v
func PrintShortestPathsTree(s, v *adjacencylist.AdjListVertex) {
	if s == v {
		fmt.Print("->", v.Value)
		fmt.Println()
		return
	}

	if v.P != nil {
		fmt.Print("->", v.Value)
		PrintShortestPathsTree(s, v.P)
	} else {
		fmt.Println(v.Value, "has no path to s, sorry!")
	}
}

func main() {
	l := buildWeightedAdjList()
	DagShortestPaths(l, s)
	fmt.Println("-=- Shortest paths from s to other vertices:")
	PrintShortestPathsTree(s, r)
	PrintShortestPathsTree(s, x)
	PrintShortestPathsTree(s, y)
	PrintShortestPathsTree(s, t)
	PrintShortestPathsTree(s, z)
}
