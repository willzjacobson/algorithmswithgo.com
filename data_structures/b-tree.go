package main

import "fmt"

// BTree : type representing a B-tree
type BTree struct {
	root BTreeNode
	t    int
}

// BTreeNode : type representing a node of a B-tree
type BTreeNode struct {
	keys     []int
	children []BTreeNode
	n        int
	leaf     bool
}

// AllocateNode : in a real implementation, this function would allocate a disk page to be used as a new node
func AllocateNode() BTreeNode {
	return BTreeNode{}
}

// BTreeCreate : creates B tree and adds an empty root
// Runs in O(1) time
func BTreeCreate(t int) BTree {
	x := AllocateNode()
	x.leaf = true
	// Here is where we'd write the new node to disk
	return BTree{
		root: x,
		t:    t,
	}
}

// BTreeSplitChild : Split a full node into 2 nodes
// We split the full node y = x.children[i] (also denoted as x.ci) about its median key S,
// which moves up into y's parent node x.
// The keys in y that are > S move into a new node z,
// which becomes a new child of x
// Runs in O(t) time, since the loops go from 0 -> t-1
func (t BTree) BTreeSplitChild(x BTreeNode, i int) {
	z := AllocateNode()
	y := x.children[i] // node to split
	z.leaf = y.leaf    // if existing node to split is a leaf, so will be its new sibling
	z.n = t.t - 1      // since y was full (and this has 2t-1 keys) and we are splitting y on S=y.keys[i], z will have t-1 keys
	// move y's keys that are > S into node z (S = y.keys[t-1], due to Go's zero indexing. So this loop goes up to j=t-2)
	for j := 0; j < t.t-1; j++ {
		z.keys[j] = y.keys[j+t.t]
	}
	// if y is not a leaf node, move its children over to node z
	if !y.leaf {
		for j := 0; j < t.t-1; j++ {
			z.children[j] = y.children[j+t.t]
		}
	}
	y.n = t.t - 1 // having moved t keys out of y, update its key count accordingly
	// scooch the higher half of x's children up by 1 to make room for the new node Z
	for j := x.n; j >= i+1; j-- {
		x.children[j+1] = x.children[j]
	}
	x.children[i+1] = z
	// scooch the higher half of x's keys up by 1 to make room for the new key S
	for j := x.n - 1; j >= i; j-- {
		x.keys[j+1] = x.keys[j]
	}
	x.keys[i] = y.keys[t.t-1] // Insert S into x's keys
	x.n = x.n + 1             // adjust x's key count to account for adding S
	// Here is where we'd perform a disck write operation for creating node z, and updating nodes x and y
}

// BTreeInsert : Insert node into an existing Btree
// splits the root if the root is full, then calls helper BTreeInsertNonFull to manage the insertion of the new key
// Runs in O(t logtn) time, since:
// - calls to BTreeSplitChild run in O(t) time
// - there will be at most logtn calls to BTreeInsertNonFull (height of BTree)
func (t BTree) BTreeInsert(k int) {
	r := t.root
	if r.n == 2*t.t-1 {
		// root node is full; split it into 2 nodes (thus increasing height of BTree by 1)
		s := AllocateNode()
		t.root = s
		s.leaf = false
		s.children = append(s.children, r)
		t.BTreeSplitChild(r, 0)
	}
	// call helper to manage the insertion of the new key. At this point, we know node r is not full.
	t.BTreeInsertNonFull(r, k)
}

// BTreeInsertNonFull : recursive function to manage insertion of a new key into the BTree
func (t BTree) BTreeInsertNonFull(x BTreeNode, k int) {
	i := x.n - 1
	if x.leaf {
		for i >= 0 && k < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = k
		x.n++
		// write updates to node x to disk
	} else {
		// find index of child node that serves as root of the subtree we need to dig into
		for i >= 0 && k < x.keys[i] {
			i--
		}
		i++
		// read x.children[i] from disk
		if x.children[i].n == 2*t.t-1 {
			// child node of x is full; must be split
			t.BTreeSplitChild(x, i)
			// if key should go in the new node, increment the key index
			if k > x.keys[i] {
				i = i + 1
			}
		}
		// keep traversing down the tree
		t.BTreeInsertNonFull(x.children[i], k)
	}
}

// BTreeSearch : given a starting node and a key k, return the node containing k and the index at which k it appears in the list of the node's keys
// If no node is found containing the key, return the leaf node checked, and -1 for the index
// time complexity is O(t logtn), since:
//  - the method could run at most logt times
//  - since x.n <= 2t-1, the for loop takes O(t) time
func (t *BTree) BTreeSearch(x BTreeNode, k int) (BTreeNode, int) {
	i := 0
	// look for the index of the lowest k in x.keys that is greater than k
	for i < x.n && k < x.keys[i] {
		i = i + 1
	}
	// if x.keys[i] is what we're looking for, return the relevant info here
	if i < x.n && x.keys[i] == k {
		return x, i
	}
	// if node x is is a leaf and we haven't found the key, return nil result
	if x.leaf {
		return x, -1
	}
	// This is where we'd have to perform a disk operation to fetch the correct child of node x
	// recursively continue down the B-tree, looking for the node that contains key k...
	return t.BTreeSearch(x.children[i], k)
}

func main() {
	t := BTreeCreate(2)
	fmt.Println("B-tree:", t)
	fmt.Println(t.root)
	n, i := t.BTreeSearch(t.root, 0)
	fmt.Println("n:", n, "i:", i)
}
