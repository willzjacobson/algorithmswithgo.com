package heap

type HeapNode struct {
	right *HeapNode
	left  *HeapNode
	value int
}

func CreateHeapNode(value int) *HeapNode {
	return &HeapNode{
		value: value,
	}
}
