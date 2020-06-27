package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"fmt"
)

const (
	white = "white"
	gray  = "gray"
	black = "black"
)

// create vertices to be referenced below
var (
	underShorts = adjacencylist.CreateAdjListVertex("underShorts")
	pants       = adjacencylist.CreateAdjListVertex("pants")
	belt        = adjacencylist.CreateAdjListVertex("belt")
	jacket      = adjacencylist.CreateAdjListVertex("jacket")
	shirt       = adjacencylist.CreateAdjListVertex("shirt")
	tie         = adjacencylist.CreateAdjListVertex("tie")
	socks       = adjacencylist.CreateAdjListVertex("socks")
	shoes       = adjacencylist.CreateAdjListVertex("shoes")
	watch       = adjacencylist.CreateAdjListVertex("watch")
)

func setUpSortTest() *adjacencylist.AdjacencyList {
	// create adjacency list
	l := adjacencylist.CreateAdjacencyList()
	// establish edges
	l.Adj[underShorts] = []*adjacencylist.AdjListVertex{pants}
	l.Adj[pants] = []*adjacencylist.AdjListVertex{belt, shoes}
	l.Adj[belt] = []*adjacencylist.AdjListVertex{jacket}
	l.Adj[shirt] = []*adjacencylist.AdjListVertex{belt, tie}
	l.Adj[tie] = []*adjacencylist.AdjListVertex{jacket}
	l.Adj[jacket] = []*adjacencylist.AdjListVertex{}
	l.Adj[shoes] = []*adjacencylist.AdjListVertex{}
	l.Adj[socks] = []*adjacencylist.AdjListVertex{shoes}
	l.Adj[watch] = []*adjacencylist.AdjListVertex{}
	return l
}

// universal time keeper for DFS
var time = 0
var sorted = []*adjacencylist.AdjListVertex{}

// TopologicalSort : Topologically sorts a directed acyclic graph represented by an adjacency-list
func TopologicalSort(g *adjacencylist.AdjacencyList) {
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
	sorted = append([]*adjacencylist.AdjListVertex{u}, sorted...)
}

func main() {
	// create adjacency-list
	l := setUpSortTest()
	TopologicalSort(l)

	// loop through sorted list
	for _, v := range sorted {
		fmt.Println(v.Value, " | start:", v.Start, " | end:", v.End)
	}
}
