package adjacencylist

// AdjListVertex : representation of a vertex in a graph, as modeled by an adjacency list
type AdjListVertex struct {
	Value string
	Color string         // used by the BFS algorithm
	D     int            // value is assigned by the BFS algorithm
	P     *AdjListVertex // used in the BFS algorithm to create a breadth-first tree
	Start int            // used as timestamp by DFS algorithm
	End   int            // used as timestamp by DFS algorithm
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
