package main

import (
	"fmt"
	"math"
)

// VEBTree : implementation of a vEB-Tree
// This implementation simplifies indexing by requiring the constraint u=2^(2^k) for any integer k.
// While this limits our set of possible values for u to 2,4,16,256,65536, it ensures u is an integer.
// If we were willing to make use of some techniques that allow u to be a non integer,
// we can have u=2^k for any integer k.
type VEBTree struct {
	u       int
	null    int
	min     int
	max     int
	summary *VEBTree
	cluster []*VEBTree
}

// CreateVEBTree : creates a new VEBTree
func CreateVEBTree(u int) *VEBTree {
	null := -1
	v := &VEBTree{
		u:    u,
		null: null,
		min:  null,
		max:  null,
	}

	if u > 2 {
		rootU := v.rootU()
		v.summary = CreateVEBTree(rootU)
		for i := 0; i < rootU; i++ {
			v.cluster = append(v.cluster, CreateVEBTree(rootU))
		}
	}

	return v
}

// Indexing Helpers
func (v *VEBTree) rootU() int {
	return int(math.Sqrt(float64(v.u)))
}

// Return the index in v.cluster that element x would reside in
func (v *VEBTree) high(x int) int {
	return x / v.rootU()
}

// Return the place in a cluster that element x would reside in
func (v *VEBTree) low(x int) int {
	return x % v.rootU()
}

// Given the index of a  cluster in v.cluster, and its index within that cluster, return the index in the full tree that element x would reside in
func (v *VEBTree) index(x, y int) int {
	return x*v.rootU() + y
}

// Minimum : Returns minimum element in a VEBTree
func (v *VEBTree) Minimum() int {
	return v.min
}

// Maximum : Returns maximum element in a VEBTree
func (v *VEBTree) Maximum() int {
	return v.max
}

// emptyTreeInsert : helper for Insert method
func (v *VEBTree) emptyTreeInsert(x int) {
	v.max = x
	v.min = x
}

// Insert : Inserts an element into a VEBTree
func (v *VEBTree) Insert(x int) {
	if v.min == v.null {
		v.emptyTreeInsert(x)
		return
	}

	if x < v.min {
		// x is the new min; we now have to place the old min into a cluster
		v.min, x = x, v.min
	}

	if v.u > 2 {
		// recursive case
		clusterForX := v.cluster[v.high(x)]
		if clusterForX.Minimum() == v.null {
			// if the cluster that would contain x is empty,
			// add it to the empty tree, and add it to the top tree's summary as well
			v.summary.Insert(v.high(x))
			clusterForX.emptyTreeInsert(v.low(x))
		} else {
			clusterForX.Insert(v.low(x))
		}
	}

	if x > v.max {
		v.max = x
	}
}

// Search : search for an element in a VEBTree.
func (v *VEBTree) Search(x int) bool {
	if v.min == x || v.max == x {
		return true
	}

	if v.u == 2 {
		// for trees where v.u=2, if it's not the min or max, it's not present
		return false
	}
	// search recursively for x in the cluster we'd expect it to be in
	return v.cluster[v.high(x)].Search(v.low(x))
}

// Successor : find an element's successor
// expects x to be in the range 0 <= x < u
func (v *VEBTree) Successor(x int) int {
	// base cases
	if v.u == 2 {
		if x == 0 && v.max == 1 {
			return 1
		}
		return v.null
	}
	if v.min != v.null && x < v.min {
		return v.min
	}

	// recursive cases
	maxElInXsCluster := v.cluster[v.high(x)].Maximum() // maximum element in x's cluster
	if maxElInXsCluster != v.null && maxElInXsCluster > v.low(x) {
		// max element in x's cluster is > x, so x's successor is in the same cluster
		// find the place of x's successor within x's cluster
		offset := v.cluster[v.high(x)].Successor(v.low(x))
		// We don't know how deep we are, so we have to recursively return the place of
		// x's successor within this tree in order to return its place within the top level tree.
		return v.index(v.high(x), offset)
	}
	// by now we know x >= the max of its cluster,
	// so we must find the next cluster that contains elements and look there for x's successor
	succCluster := v.summary.Successor(v.high(x))
	if succCluster == v.null {
		return v.null
	}
	offset := v.cluster[succCluster].Minimum()
	return v.index(succCluster, offset)
}

// Predecessor : find an element's predecessor
func (v *VEBTree) Predecessor(x int) int {
	// base cases
	if v.u == 2 {
		if v.min == 0 && x == 1 {
			return 0
		}
		return v.null
	}
	// if x is greater than the max of the tree, just return the max of the tree
	if v.max != v.null && x > v.max {
		return v.max
	}

	// recursive cases
	minElInXsCluster := v.cluster[v.high(x)].Minimum()
	if minElInXsCluster != v.null && minElInXsCluster < v.low(x) {
		// an element smaller than x (and thus the predecessor) exists in x's cluster
		offset := v.cluster[v.high(x)].Predecessor(v.low(x))
		return v.index(v.high(x), offset)
	}
	// x's predeccessor is in a different cluster. Use the tree's summary to find that cluster and dig into it
	predCluster := v.summary.Predecessor(v.high(x))
	if predCluster == v.null {
		if v.min != v.null && v.min < x {
			return v.min
		}
		return v.null
	}
	offset := v.cluster[predCluster].Maximum()
	return v.index(predCluster, offset)
}

// Delete : Delete an element from a VEBTree
func (v *VEBTree) Delete(x int) {
	// base cases
	if v.min == v.max {
		// v contains only 1 element
		v.min = v.null
		v.max = v.null
		return
	}
	if v.u == 2 {
		if x == 1 {
			v.max = 0
		} else {
			v.min = 1
		}
		return
	}

	// recursive case: v has >1 element, and v.u>=4
	if x == v.min {
		firstCluster := v.summary.Minimum() // first cluster containing an element (v.min is not an a cluster)
		// set x to the value of the lowest element in that cluser (the new v.min, since we're deleting x)
		x = v.index(firstCluster, v.cluster[firstCluster].Minimum())
		v.min = x
		// having reassigned the variable x to a different element in the cluster and made that element the new min,
		// we must now delete x from the cluster (since the v.min is not in a cluster)
	}
	// we must delete x from its cluster, whether the variable x has been reassigned or not
	v.cluster[v.high(x)].Delete(v.low(x))

	// x's cluster may now be empty. If it is, we must remove x's cluster from the summary
	if v.cluster[v.high(x)].Minimum() == v.null {
		v.summary.Delete(v.high(x))
		if x == v.max {
			// check whether x was the summary max, and if it was, we need to update the summary max to the highest remaining element in the summary
			summaryMax := v.summary.Maximum() // the highest numbered nonempty cluster
			if summaryMax == v.null {
				// All clusters are empty, so only the v.min remains in V.
				v.max = v.min
			} else {
				// set max to the maximum element in the highest-numbered nonempty cluster
				v.max = v.index(summaryMax, v.cluster[summaryMax].Maximum())
			}
		}
	} else if x == v.max {
		// x's cluster did not become empty when x was deleted
		v.max = v.index(v.high(x), v.cluster[v.high(x)].Maximum())
	}
}

func main() {
	v := CreateVEBTree(16)
	fmt.Println("VEB Tree for u=16:", v)
	fmt.Println("VEB Tree summary:", v.summary)
	fmt.Println("VEB Tree first cluster:", v.cluster[0])
	fmt.Println("VEB Tree first cluster's first cluster:", v.cluster[0].cluster[0])
	fmt.Println("-=-=-=- Inserting now...")
	v.Insert(2)
	v.Insert(3)
	v.Insert(4)
	v.Insert(5)
	v.Insert(7)
	v.Insert(14)
	fmt.Println("VEB Tree 4th cluster:", v.cluster[3])
	fmt.Println("VEB Tree 4th cluster's 2nd cluster:", v.cluster[3].cluster[1].min, v.cluster[3].cluster[1].max)
	v.Insert(15)
	fmt.Println("VEB Tree 4th cluster:", v.cluster[3])
	fmt.Println("VEB Tree 4th cluster's 2nd cluster:", v.cluster[3].cluster[1].min, v.cluster[3].cluster[1].max)
	fmt.Println("-=-=-=- Done Inserting...")
	fmt.Println("VEB Tree for u=16:", v)
	fmt.Println("VEB Tree summary:", v.summary)
	fmt.Println("VEB Tree first cluster:", v.cluster[0])
	fmt.Println("VEB Tree first cluster's first cluster:", v.cluster[0].cluster[0])
	fmt.Println("-=-=- Searching...")
	fmt.Println("for 2 (present):", v.Search(2))
	fmt.Println("for 15 (present):", v.Search(15))
	fmt.Println("for 4 (present):", v.Search(4))
	fmt.Println("for 6 (NOT present):", v.Search(6))
	fmt.Println("-=-=- Successor...")
	fmt.Println("0:", v.Successor(0))
	fmt.Println("2:", v.Successor(2))
	fmt.Println("3:", v.Successor(3))
	fmt.Println("4:", v.Successor(4))
	fmt.Println("6:", v.Successor(6))
	fmt.Println("7:", v.Successor(7))
	fmt.Println("13:", v.Successor(13))
	fmt.Println("15:", v.Successor(15))
	fmt.Println("-=-=- Predecessor...")
	fmt.Println("0:", v.Predecessor(0))
	fmt.Println("2:", v.Predecessor(2))
	fmt.Println("3:", v.Predecessor(3))
	fmt.Println("4:", v.Predecessor(4))
	fmt.Println("6:", v.Predecessor(6))
	fmt.Println("7:", v.Predecessor(7))
	fmt.Println("13:", v.Predecessor(13))
	fmt.Println("15:", v.Predecessor(15))
	fmt.Println("-=-=- Delete...")
	fmt.Println("VEB Tree 4th cluster:", v.cluster[3])
	fmt.Println("VEB Tree 4th cluster's 2nd cluster:", v.cluster[3].cluster[1].min, v.cluster[3].cluster[1].max)
	v.Delete(15)
	fmt.Println("VEB Tree 4th cluster:", v.cluster[3])
	fmt.Println("VEB Tree 4th cluster's 2nd cluster:", v.cluster[3].cluster[1].min, v.cluster[3].cluster[1].max)
}
