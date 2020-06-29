package adjacencylist

// AdjListVertex : representation of a vertex in a graph, as modeled by an adjacency list
type AdjListVertex struct {
	Value string
	Color string         // used by the BFS algorithm
	D     int            // value is assigned by the BFS algorithm
	P     *AdjListVertex // used in the BFS algorithm to create a breadth-first tree
	Start int            // used as timestamp by DFS algorithm
	End   int            // used as timestamp by DFS algorithm
	rank  int            // used to manage disjoint sets of vertices in Kruskal's algo for determining a minimum spanning tree (MST) for a graph
	Key   float64        // used in prim's MST algo to manage the min-heap
	Index int            // used in prim's MST algo to manage the min-heap (to provide index for DecreaseKeys method)
}

// CreateAdjListVertex : creates a new Vertex
func CreateAdjListVertex(v string) *AdjListVertex {
	return &AdjListVertex{
		Value: v,
	}
}

// AdjacencyList : model of an adjacency-list
// contains the data structure as "Adj", which is a map
// where each kay is a pointer to a vertex, and its value is
// an array of pointers to other vertices with wwhich it shares an edge
type AdjacencyList struct {
	Adj       map[*AdjListVertex][]*AdjListVertex
	Transpose map[*AdjListVertex][]*AdjListVertex
}

// CreateAdjacencyList : creates an empty AdjacencyList
func CreateAdjacencyList() *AdjacencyList {
	m := make(map[*AdjListVertex][]*AdjListVertex)
	return &AdjacencyList{
		Adj: m,
	}
}

// TransposeAdjacencyList : builds a new adjacency-list to hold the transpose of the original
// Only relevant for directed graphs
// Runtime: O(V+E)
func (l *AdjacencyList) TransposeAdjacencyList() {
	// create keys for each vertex
	for k := range l.Adj {
		l.Transpose[k] = []*AdjListVertex{}
	}

	// create the edges for each vertex
	for k, descedants := range l.Adj {
		for _, v := range descedants {
			l.Transpose[v] = append(l.Transpose[v], k)
		}
	}
}

// ** Weighted implementation

// Weighted : Adjacency List representation for a graph with weighted edges
type Weighted struct {
	Adj       map[*AdjListVertex][]*AdjListEdgeWeighted
	Transpose map[*AdjListVertex][]*AdjListEdgeWeighted
}

// CreateWeighted : creates an empty AdjacencyList
func CreateWeighted() *Weighted {
	m := make(map[*AdjListVertex][]*AdjListEdgeWeighted)
	return &Weighted{
		Adj: m,
	}
}

// AdjListEdgeWeighted : for storage of a weighted edge in an adjacency-list
type AdjListEdgeWeighted struct {
	To     *AdjListVertex
	Weight float64
}

// CreateAdjListEdgeWeighted : create a AdjListEdgeWeighted
func CreateAdjListEdgeWeighted(v *AdjListVertex, w float64) *AdjListEdgeWeighted {
	return &AdjListEdgeWeighted{
		To:     v,
		Weight: w,
	}
}

// Edge utility (sortable)

// Edge : type used in implementation of Kruskal's algo
// the algo requires sorting all the edges in a graph in order of weight.
// This type is not required in a standard adjacency list representation of a graph
type Edge struct {
	From   *AdjListVertex
	To     *AdjListVertex
	Weight float64
}

// EdgesByWeight : sortable slice of Edges
type EdgesByWeight []*Edge

func (w EdgesByWeight) Len() int {
	return len(w)
}
func (w EdgesByWeight) Less(i, j int) bool {
	return w[i].Weight < w[j].Weight
}
func (w EdgesByWeight) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

// **  Disjoint Set Management

// Set : representation of a disjoint set of vertices
type Set struct {
	head *AdjListVertex
}

// CreateDisjointSet : create a disjoint set of vertices
func CreateDisjointSet(x *AdjListVertex) *Set {
	x.rank = 0
	x.P = x
	return &Set{
		head: x,
	}
}

// Union : join 2 disjoint sets of vertices
func Union(x *AdjListVertex, y *AdjListVertex) {
	Link(FindSet(x), FindSet(y))
}

// Link : helper used by Union method to join 2 disjoint sets of vertices
func Link(x *AdjListVertex, y *AdjListVertex) {
	if x.rank > y.rank {
		y.P = x
	} else {
		x.P = y
		if x.rank == y.rank {
			y.rank++
		}
	}
}

// FindSet : return the head of a disjoint set of graph vertices
// in the process of traversing up the tree representing the set,
// for each node x encountered, set x.P -> the head of the tree
func FindSet(x *AdjListVertex) *AdjListVertex {
	if x.P != x {
		x.P = FindSet(x.P)
	}
	return x.P
}
