package heap

/*
MaxHeap :
A heap is represented by an array A that has 2 attributes:
	A.length, which gives the number of elements in the array
	a.heapSize, which represents jw many elements in hte heap are stored within A
While A[1...A.length] may contain numbers,
only the elements A[1...A.heapSize], where 1 <= A.heapSize <= A.length
are valid members of the heap.

The root of the tree is A[1]
We can compute the indices of of an element's parent, left, and right child (in this example we index at 1, not 0):
	parent element is i/2
	left child is 2i
	right child is 2i+1

The subarray A[(n/2)+1...n] are all leaves of the tree (and thus are each a 1-element heap)
Thus, when converting an array to a heap, we need only address the elements 1...n/2

Basic heap operations run in time with the height of the tree representation: O(log n)
*/
type MaxHeap struct {
	Size  int
	Slice []int
}

func CreateMaxHeap() *MaxHeap {
	return &MaxHeap{
		Slice: []int{},
	}
}

// Only need to worry about the nodes that are not leaves
func (h *MaxHeap) AssignList(nums *[]int) {
	if h.Size > 0 {
		panic("This method is only intended for when the heap is empty")
	}
	h.Slice = *nums
	h.Size = len(*nums)
	for i := len(h.Slice) / 2; i >= 0; i-- {
		h.SendNodeDown(i)
	}
}

func (h *MaxHeap) Insert(val int) {
	h.Slice = append(h.Slice, val)
	h.Size++
	h.bringNodeUp(h.Size - 1)
}

func (h *MaxHeap) Maximum() int {
	if h.Size < 1 {
		panic("heap underflow")
	}
	return h.Slice[0]
}

func (h *MaxHeap) ExtactMax() int {
	if h.Size < 1 {
		panic("heap underflow")
	}
	max := h.Slice[0]
	h.Slice[0] = h.Slice[h.Size-1]
	h.Size--
	h.SendNodeDown(0)
	return max
}

func (h *MaxHeap) IncreaseKeys(i int, increaseTo int) {
	if h.Size <= i {
		panic("heap underflow")
	}
	h.Slice[i] = increaseTo
	h.bringNodeUp(i)
}

func (h *MaxHeap) bringNodeUp(i int) {
	parentInd := i / 2
	for h.Slice[i] > h.Slice[parentInd] {
		h.Slice[i], h.Slice[parentInd] = h.Slice[parentInd], h.Slice[i]
		i = parentInd
		parentInd = i / 2
	}
}

func (h *MaxHeap) SendNodeDown(i int) {
	largestInd := i
	largest := h.Slice[i]
	leftInd := i * 2
	rightInd := i*2 + 1

	if h.Size >= leftInd+1 && h.Slice[leftInd] > largest {
		largest = h.Slice[leftInd]
		largestInd = leftInd
	}
	if h.Size >= rightInd+1 && h.Slice[rightInd] > largest {
		largest = h.Slice[leftInd]
		largestInd = rightInd
	}

	if largestInd != i {
		h.Slice[largestInd], h.Slice[i] = h.Slice[i], h.Slice[largestInd]
		h.SendNodeDown(largestInd)
	}
}
