package main

import (
	adjacencylist "algo/graph-algorithms/adjacency-list"
	minheap "algo/graph-algorithms/min-heap"
	"fmt"
	"math"
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

func buildWeightedAdjListForPrim() *adjacencylist.Weighted {
	// create adjacency list
	l := adjacencylist.CreateWeighted()
	// establish edges (important to get both directions)
	l.Adj[a] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(b, 4),
		adjacencylist.CreateAdjListEdgeWeighted(h, 8),
	}
	l.Adj[b] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(a, 4),
		adjacencylist.CreateAdjListEdgeWeighted(c, 8),
		adjacencylist.CreateAdjListEdgeWeighted(h, 11),
	}
	l.Adj[c] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(b, 8),
		adjacencylist.CreateAdjListEdgeWeighted(d, 7),
		adjacencylist.CreateAdjListEdgeWeighted(f, 4),
		adjacencylist.CreateAdjListEdgeWeighted(i, 2),
	}
	l.Adj[d] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(c, 7),
		adjacencylist.CreateAdjListEdgeWeighted(e, 9),
		adjacencylist.CreateAdjListEdgeWeighted(f, 14),
	}
	l.Adj[e] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(d, 9),
		adjacencylist.CreateAdjListEdgeWeighted(f, 10),
	}
	l.Adj[f] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(g, 2),
		adjacencylist.CreateAdjListEdgeWeighted(e, 10),
		adjacencylist.CreateAdjListEdgeWeighted(c, 4),
		adjacencylist.CreateAdjListEdgeWeighted(g, 2),
	}
	l.Adj[g] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(f, 2),
		adjacencylist.CreateAdjListEdgeWeighted(h, 1),
		adjacencylist.CreateAdjListEdgeWeighted(i, 6),
	}
	l.Adj[h] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(a, 8),
		adjacencylist.CreateAdjListEdgeWeighted(g, 1),
		adjacencylist.CreateAdjListEdgeWeighted(i, 7),
	}
	l.Adj[i] = []*adjacencylist.AdjListEdgeWeighted{
		adjacencylist.CreateAdjListEdgeWeighted(c, 2),
		adjacencylist.CreateAdjListEdgeWeighted(g, 6),
		adjacencylist.CreateAdjListEdgeWeighted(h, 7),
	}
	return l
}

// MSTPrim : implementation of Prim's algorithm for determining a minimum spanning tree of a graph
// updates the parent prop of each vertex to point to the vertex via which it joined the tree
// 2nd parameter is an arbitrary vertex to serve as the root for the MST
// returns the total weight, just so we can check our answer
func MSTPrim(l *adjacencylist.Weighted, r *adjacencylist.AdjListVertex) float64 {
	var totalWeight float64 = 0                                           // keep track of the total weight of the MST
	q := minheap.CreateMinHeap()                                          // create a min heap to manage the greediness
	dictOfVerticesStillInQueue := map[*adjacencylist.AdjListVertex]bool{} // data structure to keep track of which vertices are still in the queue (min heap)
	r.Key = 0                                                             // to ensure it is dequeued from the min-heap first
	// set the Key of each vertex to infinity (except the root), and add them to the queue
	for v := range l.Adj {
		if v != r {
			v.Key = math.Inf(1) // initially, set keys to +infinity
		}
		q.Insert(v) // enqueue each vertex in the min heap
		dictOfVerticesStillInQueue[v] = true
	}

	for q.Size > 0 {
		// the next vertex extracted from the min-heap will not yet be in the tree,
		// but will have the least weight to be added to the tree compared to all other vertices not yet in the tree.
		// Hence it will be the greedy choice, and thus the next one to attach
		u := q.ExtractMin()
		totalWeight += u.Key                  // add the cost incurred to add this vertex to the tree
		dictOfVerticesStillInQueue[u] = false // document that this vertex is no longer in the queue
		for _, e := range l.Adj[u] {          // loop through all the edges that this vertex has with other vertices
			// since we've just added u to the tree, we can now update the keys of the adjacent vertices to reflect
			// the cost of adding each one to the tree by making u its parent
			if dictOfVerticesStillInQueue[e.To] && e.Weight < e.To.Key {
				e.To.P = u
				q.DecreaseKey(e.To.Index, e.Weight) // DecreaseKey call may reorder the vertex in the queue, enforcing our greediness
			}
		}
	}

	return totalWeight
}

func main() {
	l := buildWeightedAdjListForPrim()
	w := MSTPrim(l, a)
	fmt.Println("-=-=-")
	fmt.Println("Total weight:", w)
	fmt.Println("i's parent:", i.P) // log out a couple parent-child links, just as a sanity check
	fmt.Println("b's parent:", b.P)
}
