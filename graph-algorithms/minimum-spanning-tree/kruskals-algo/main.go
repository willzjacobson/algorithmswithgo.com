package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	"fmt"
	"sort"
)

// create vertices to be referenced below
var (
	// for Krushal's algo demo
	a = adjacencylist.CreateAdjListVertex("a")
	b = adjacencylist.CreateAdjListVertex("b")
	c = adjacencylist.CreateAdjListVertex("c")
	d = adjacencylist.CreateAdjListVertex("d")
	e = adjacencylist.CreateAdjListVertex("e")
	f = adjacencylist.CreateAdjListVertex("f")
	g = adjacencylist.CreateAdjListVertex("g")
	h = adjacencylist.CreateAdjListVertex("h")
	i = adjacencylist.CreateAdjListVertex("i")
)

func buildWeightedAdjListForKruskal() *adjacencylist.Weighted {
	// create adjacency list
	l := adjacencylist.CreateWeighted()
	// establish edges
	l.Adj[a] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(b, 4), adjacencylist.CreateAdjListEdgeWeighted(h, 8)}
	l.Adj[b] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(c, 8), adjacencylist.CreateAdjListEdgeWeighted(h, 11)}
	l.Adj[c] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(d, 7), adjacencylist.CreateAdjListEdgeWeighted(f, 4), adjacencylist.CreateAdjListEdgeWeighted(i, 2)}
	l.Adj[d] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(e, 9), adjacencylist.CreateAdjListEdgeWeighted(f, 14)}
	l.Adj[e] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(f, 10)}
	l.Adj[f] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(g, 2)}
	l.Adj[g] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(h, 1), adjacencylist.CreateAdjListEdgeWeighted(i, 6)}
	l.Adj[h] = []*adjacencylist.AdjListEdgeWeighted{adjacencylist.CreateAdjListEdgeWeighted(i, 7)}
	l.Adj[i] = []*adjacencylist.AdjListEdgeWeighted{}
	return l
}

// MSTKruskal : implementation of Kruskal's algorithm for determining a minimum spanning tree of a graph
// returns a list of edges included in the resulting MST, and the total weight of the edges
func MSTKruskal(l *adjacencylist.Weighted) ([]*adjacencylist.Edge, float64) {
	edges := adjacencylist.EdgesByWeight{} // will hold all edges in the graph
	var totalWeight float64 = 0            // total weight of the resulting MST
	a := adjacencylist.EdgesByWeight{}     // will hold all edges in the MST

	// Build a structure to keep track of disjoint sets of vertices in the graph as we build an MST
	// This structure will allow us to determine whether any 2 vertices are currently connected
	disjointSets := map[*adjacencylist.AdjListVertex]*adjacencylist.Set{}
	for v := range l.Adj {
		disjointSets[v] = adjacencylist.CreateDisjointSet(v)
	}

	// Build array of all egdes the in the graph
	for v := range l.Adj {
		for _, e := range l.Adj[v] {
			edge := &adjacencylist.Edge{
				From:   v,
				To:     e.To,
				Weight: e.Weight,
			}
			edges = append(edges, edge)
		}
	}
	// Sort edges in non-decreasing order by weight
	sort.Sort(edges)

	// Aggregate edges to be included in the MST by greedily choosing
	// the lowest weight edge that connects 2 disjoint sets of edges (that is, a light edge)
	for _, e := range edges {
		if adjacencylist.FindSet(e.To) != adjacencylist.FindSet(e.From) {
			// These 2 vertices are not yet connected. We must:
			// - connect them
			// - add the weight of the edge to the total weight of the MST
			// - add the edge to the MST
			adjacencylist.Union(e.To, e.From)
			totalWeight += e.Weight
			a = append(a, e)
		}
	}
	// Return the edges in the MST and their total weight
	return a, totalWeight
}

func main() {
	l := buildWeightedAdjListForKruskal()
	edges, totalWeight := MSTKruskal(l)
	fmt.Println("--")
	fmt.Println("Total Weight of MST", totalWeight)
	fmt.Println("The", len(edges), "edges included in the MST:")
	fmt.Println("-=-=-")
	for _, e := range edges {
		fmt.Println("From:", e.From.Value, "To", e.To.Value, "Weight:", e.Weight)
	}
}
