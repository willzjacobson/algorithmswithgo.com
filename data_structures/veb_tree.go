package main

import (
	"fmt"
	"math"
)

// VEBTree : implementation of a vEB-Tree
// This implementation simplifies indexing by requiring the constraint u=2^(2^2k) for any integer k.
// While this limits our set of possible values for u to 2,3,16,256,65536, it ensures u is an integer.
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
	v := &VEBTree{
		u:    u,
		null: -1,
		min:  -1,
		max:  -1,
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
func (v *VEBTree) Successor(x int) int {
	return -1
}

// Predecessor : find an element's predecessor
func (v *VEBTree) Predecessor(x int) int {
	return -1
}

// Delete : Delete an element from a VEBTree
func (v *VEBTree) Delete(x int) {

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
	v.Insert(15)
	fmt.Println("VEB Tree for u=16:", v)
	fmt.Println("VEB Tree summary:", v.summary)
	fmt.Println("VEB Tree first cluster:", v.cluster[0])
	fmt.Println("VEB Tree first cluster's first cluster:", v.cluster[0].cluster[0])
	fmt.Println("-=-=- Searching...")
	fmt.Println("for 2 (present):", v.Search(2))
	fmt.Println("for 15 (present):", v.Search(15))
	fmt.Println("for 4 (present):", v.Search(4))
	fmt.Println("for 6 (NOT present):", v.Search(6))
}
