package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"algo/graph-algorithms/all-pairs-shortest-paths/helpers"
	bf "algo/graph-algorithms/bellman-ford"
	"algo/graph-algorithms/dijkstra"
	"fmt"
	"math"
	"strconv"
)

type extendedAdjListVertex struct {
}

// create vertices to be referenced below
var (
	zero  = adjacencylist.CreateAdjListVertex("0")
	one   = adjacencylist.CreateAdjListVertex("1")
	two   = adjacencylist.CreateAdjListVertex("2")
	three = adjacencylist.CreateAdjListVertex("3")
	four  = adjacencylist.CreateAdjListVertex("4")
)

func buildWeightedAdjList() *adjacencylist.Weighted {
	// create adjacency list
	l := adjacencylist.CreateWeighted()
	// establish edges
	l.Adj[zero] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(one, 3),
		adjacencylist.CreateAdjListEdgeWeighted(two, 8),
		adjacencylist.CreateAdjListEdgeWeighted(four, -4),
	}
	l.Adj[one] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(three, 1),
		adjacencylist.CreateAdjListEdgeWeighted(four, 7),
	}
	l.Adj[two] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(one, 4),
	}
	l.Adj[three] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(zero, 2),
		adjacencylist.CreateAdjListEdgeWeighted(two, -5),
	}
	l.Adj[four] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(three, 6),
	}
	return l
}

// helpers
func generateNewAdjListWithS(l *adjacencylist.Weighted, s *adjacencylist.AdjListVertex) *adjacencylist.Weighted {
	// generate a new graph that will be an extension of l
	lPrime := adjacencylist.CreateWeighted()
	// add a new adjacency-list for a new vertex s
	lPrime.Adj[s] = []*adjacencylist.AdjListEdgeWeighted{}
	// add adjacency-lists from l to lPrime, and establish a 0-weight edge from s to all other vertices
	for v := range l.Adj {
		lPrime.Adj[v] = l.Adj[v]
		lPrime.Adj[s] = append(lPrime.Adj[s], adjacencylist.CreateAdjListEdgeWeighted(v, 0))
	}
	// return new extended graph
	return lPrime
}

// Johnson : implementation of Johnson's all-pairs shortest-paths algorithm
func Johnson(l *adjacencylist.Weighted) helpers.AdjacencyMatrix {
	// create empty matrix to become shortest-paths matrix
	d := helpers.GenerateNewMatrix(len(l.Adj), math.Inf(1))
	// create theoretical source vertex s to use for negative-weight cycle check with Bellman-Ford algo
	s := &adjacencylist.AdjListVertex{}
	lPrime := generateNewAdjListWithS(l, s)
	containsNegWeightCycles := !bf.BellmanFordSSSP(lPrime, s)

	if containsNegWeightCycles {
		fmt.Println("Graph contains negative weight cycles: cannot proceed to compute shortest paths")
		return d
	}

	// fmt.Println("from s to 0", zero.D)
	// fmt.Println("from s to 1", one.D)
	// fmt.Println("from s to 2", two.D)
	// fmt.Println("from s to 3", three.D)
	// fmt.Println("from s to 4", four.D)

	// For each vertex, Set v.H to be the shortest path from 𝛿(s,v)
	for u := range lPrime.Adj {
		u.H = u.D
	}

	// For each edge (u,v), set the new weight w' (w'(u,v) = w(u,v) + u.H - v.H)
	// (cannot leverage the outer loop above, since we want v.H values set for all v in V)
	for u := range l.Adj {
		for _, e := range l.Adj[u] {
			e.Weight = e.Weight + u.D - e.To.D
		}
	}

	// fmt.Println("-=-=- test reweighting for 0")
	// for _, e := range l.Adj[zero] {
	// 	fmt.Println("To:", e.To.Value, "weight:", e.Weight)
	// }
	// fmt.Println("-=-=- test reweighting for 1")
	// for _, e := range l.Adj[one] {
	// 	fmt.Println("To:", e.To.Value, "weight:", e.Weight)
	// }
	// fmt.Println("-=-=- test reweighting for 2")
	// for _, e := range l.Adj[two] {
	// 	fmt.Println("To:", e.To.Value, "weight:", e.Weight)
	// }
	// fmt.Println("-=-=- test reweighting for 3")
	// for _, e := range l.Adj[three] {
	// 	fmt.Println("To:", e.To.Value, "weight:", e.Weight)
	// }
	// fmt.Println("-=-=- test reweighting for 4")
	// for _, e := range l.Adj[four] {
	// 	fmt.Println("To:", e.To.Value, "weight:", e.Weight)
	// }
	c := 0

	for u := range l.Adj {
		c++
		dijkstra.Dijkstra(l, u)
		fmt.Println("-=-=-")
		fmt.Println("paths for", u.Value, ":")
		PrintShortestPath(u, zero)
		fmt.Println("distance to 0:", zero.D)
		PrintShortestPath(u, one)
		fmt.Println("distance to 1:", one.D)
		PrintShortestPath(u, two)
		fmt.Println("distance to 2:", two.D)
		PrintShortestPath(u, three)
		fmt.Println("distance to 3:", three.D)
		PrintShortestPath(u, four)
		fmt.Println("distance to 4:", four.D)
		for v := range l.Adj {
			uInt, _ := strconv.Atoi(u.Value)
			vInt, _ := strconv.Atoi(v.Value)
			d[uInt][vInt] = u.D + v.H - u.H
		}
		if c > 1 {
			break
		}
	}

	return d
}

// PrintShortestPath : helper to show the shortest path from the source vertex to any vertex v
func PrintShortestPath(s *adjacencylist.AdjListVertex, v *adjacencylist.AdjListVertex) {
	if v == s {
		fmt.Print("->", v.Value)
		fmt.Println()
	} else if v.P != nil {
		fmt.Print("->", v.Value)
		PrintShortestPath(s, v.P)
	} else {
		fmt.Print(" No path from source vertex s to vertex", v.Value)
		fmt.Println()
	}
}

func main() {
	l := buildWeightedAdjList()
	d := Johnson(l)
	fmt.Println("-=- Shortest paths matrix generated from Johnson's algorithm:")
	fmt.Println(d)
}
