package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"algo/graph-algorithms/bfs"
	"algo/graph-algorithms/dfs"
	"fmt"
)

// create vertices to be referenced below
var (
	// for BFS demo
	r = adjacencylist.CreateAdjListVertex("r")
	s = adjacencylist.CreateAdjListVertex("s")
	t = adjacencylist.CreateAdjListVertex("t")
	u = adjacencylist.CreateAdjListVertex("u")
	v = adjacencylist.CreateAdjListVertex("v")
	w = adjacencylist.CreateAdjListVertex("w")
	x = adjacencylist.CreateAdjListVertex("x")
	y = adjacencylist.CreateAdjListVertex("y")
	// for DFS demo
	a  = adjacencylist.CreateAdjListVertex("a")
	b  = adjacencylist.CreateAdjListVertex("b")
	c  = adjacencylist.CreateAdjListVertex("c")
	d  = adjacencylist.CreateAdjListVertex("d")
	e  = adjacencylist.CreateAdjListVertex("e")
	f  = adjacencylist.CreateAdjListVertex("f")
	gg = adjacencylist.CreateAdjListVertex("g")
	h  = adjacencylist.CreateAdjListVertex("h")
)

func buildTestAdjListForBFS() *adjacencylist.AdjacencyList {
	// create adjacency list
	l := adjacencylist.CreateAdjacencyList()
	// establish edges
	l.Adj[r] = []*adjacencylist.AdjListVertex{s, v}
	l.Adj[s] = []*adjacencylist.AdjListVertex{r, w}
	l.Adj[t] = []*adjacencylist.AdjListVertex{w, x, u}
	l.Adj[u] = []*adjacencylist.AdjListVertex{t, x, y}
	l.Adj[v] = []*adjacencylist.AdjListVertex{r}
	l.Adj[w] = []*adjacencylist.AdjListVertex{s, t, x}
	l.Adj[x] = []*adjacencylist.AdjListVertex{w, t, u, y}
	l.Adj[y] = []*adjacencylist.AdjListVertex{x, u}
	return l
}

func buildTestAdjListForDFS() *adjacencylist.AdjacencyList {
	// create adjacency list
	l := adjacencylist.CreateAdjacencyList()
	// establish edges
	l.Adj[a] = []*adjacencylist.AdjListVertex{e}
	l.Adj[b] = []*adjacencylist.AdjListVertex{a, f}
	l.Adj[c] = []*adjacencylist.AdjListVertex{b, f}
	l.Adj[d] = []*adjacencylist.AdjListVertex{gg, h}
	l.Adj[e] = []*adjacencylist.AdjListVertex{b}
	l.Adj[f] = []*adjacencylist.AdjListVertex{e}
	l.Adj[gg] = []*adjacencylist.AdjListVertex{c, f}
	l.Adj[h] = []*adjacencylist.AdjListVertex{gg, d}
	return l
}

func main() {
	// Breadth First Search demo
	fmt.Println("-=-=- BFS")
	g := buildTestAdjListForBFS()
	bfs.BFS(g, s)
	fmt.Println("path from s to s:")
	bfs.PrintTree(g, s, s)
	fmt.Println("-=-")
	fmt.Println("path from s to t:")
	bfs.PrintTree(g, s, t)
	fmt.Println("-=-")
	fmt.Println("path from s to y:")
	bfs.PrintTree(g, s, y)
	fmt.Println("-=-")
	fmt.Println("path from s to u:")
	bfs.PrintTree(g, s, u)
	fmt.Println("-=-")
	z := adjacencylist.CreateAdjListVertex("z") // unconnected vertex
	fmt.Println("path from s to z (does not exist):")
	bfs.PrintTree(g, s, z)

	// Depth First Search demo
	fmt.Println("\n-=-=- DFS")
	g2 := buildTestAdjListForDFS()
	dfs.DFS(g2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(gg)
	fmt.Println(h)
}
