package main

import "fmt"

// BTree : type representing a B-tree
type BTree struct {
	root     *BTreeNode
	t        int
	nullNode *BTreeNode
}

// BTreeNode : type representing a node of a B-tree
type BTreeNode struct {
	keys     []int
	children []*BTreeNode
	n        int
	leaf     bool
}

// AllocateNode : in a real implementation, this function would allocate a disk page to be used as a new node
func (t *BTree) AllocateNode() *BTreeNode {
	x := BTreeNode{}
	for i := 0; i < 2*t.t; i++ {
		x.children = append(x.children, t.nullNode)
	}
	for i := 0; i < 2*t.t-1; i++ {
		x.keys = append(x.keys, -1)
	}
	return &x
}

// BTreeCreate : creates B tree and adds an empty root
// Runs in O(1) time
func BTreeCreate(t int) *BTree {
	// create null node to use as place filler
	nullNode := &BTreeNode{}
	for i := 0; i < 2*t; i++ {
		nullNode.children = append(nullNode.children, nullNode)
	}

	// create the tree
	tree := BTree{
		t:        t,
		nullNode: nullNode,
	}

	// create root node
	x := tree.AllocateNode()
	x.leaf = true
	tree.root = x

	// create null node used to auto-populate children of newly allocated nodes
	tree.nullNode = tree.AllocateNode()

	// *Here is where we'd write the new node to disk
	return &tree
}

// BTreeSplitChild : Split a full node into 2 nodes
// We split the full node y = x.children[i] (also denoted as x.ci) about its median key S,
// which moves up into y's parent node x.
// The keys in y that are > S move into a new node z,
// which becomes a new child of x
// Runs in O(t) time, since the loops go from 0 -> t-1
func (t *BTree) BTreeSplitChild(x *BTreeNode, i int) {
	z := t.AllocateNode()
	y := x.children[i] // y is node to split
	z.leaf = y.leaf    // if existing node to split is a leaf, so will be its new sibling
	z.n = t.t - 1      // since y was full (and this has 2t-1 keys) and we are splitting y on S=y.keys[i], z will have t-1 keys

	// move y's keys that are > S into node z (S = y.keys[t-1], due to Go's zero indexing. So this loop goes up to j=t-2)
	for j := 0; j < t.t-1; j++ {
		z.keys[j] = y.keys[j+t.t]
		y.keys[j+t.t] = -1
	}
	// if y is not a leaf node, move its children over to node z
	if !y.leaf {
		for j := 0; j < t.t-1; j++ {
			z.children[j] = y.children[j+t.t]
			y.children[j+t.t] = t.nullNode
		}
	}

	y.n = t.t // having moved t-1 keys out of y, update its key count accordingly

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
	y.keys[t.t-1] = -1        // nullify y.keys[t]
	y.n--                     // lower y.n to account for nullifying y.keys[t]
	x.n++                     // adjust x's key count to account for adding S
	// *Here is where we'd perform a disk write operation for creating node z, and updating nodes x and y
}

// BTreeInsert : Insert node into an existing Btree
// splits the root if the root is full, then calls helper BTreeInsertNonFull to manage the insertion of the new key
// Runs in O(t logtn) time, since:
// - calls to BTreeSplitChild run in O(t) time
// - there will be at most logtn calls to BTreeInsertNonFull (height of BTree)
func (t *BTree) BTreeInsert(k int) {
	r := t.root
	if r.n == 2*t.t-1 {
		// root node is full; split it into 2 nodes (thus increasing height of BTree by 1)
		s := t.AllocateNode()
		s.children[0] = r
		s.leaf = false
		t.root = s
		t.BTreeSplitChild(s, 0)
		// call helper to manage the insertion of the new key. At this point, we know node r is not full.
		t.BTreeInsertNonFull(s, k)
	} else {
		t.BTreeInsertNonFull(r, k)
	}
}

// BTreeInsertNonFull : recursive function to manage insertion of a new key into the BTree
func (t *BTree) BTreeInsertNonFull(x *BTreeNode, k int) {
	i := x.n
	if x.leaf {
		for x.n > 0 && i >= 1 && (x.keys[i-1] == -1 || k < x.keys[i-1]) {
			i--
		}

		// scooch the keys of x that are greater than k up 1 index to make room for k
		for j := 2*t.t - 3; j >= i; j-- {
			x.keys[j+1] = x.keys[j]
		}
		x.keys[i] = k // assign new key k to proper index of x.keys
		x.n++         // having added the key to the node, increment the key count
		// write updates to node x to disk
	} else {
		// find index of child node that serves as root of the subtree we need to dig into
		for i >= 0 && (x.keys[i] == -1 || k < x.keys[i]) {
			i--
		}
		i++
		// *Here is where we'd read x.children[i] from disk
		// child node of x is full; split it
		if x.children[i].n == 2*t.t-1 {
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
func (t *BTree) BTreeSearch(x *BTreeNode, k int) (*BTreeNode, int) {
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
	fmt.Println("Empty B-tree:", t.root.keys, t.root.children)
	fmt.Println(t.root)

	n1, i1 := t.BTreeSearch(t.root, 13)
	fmt.Println("Key not found:", "n:", n1, "i:", i1)

	t.BTreeInsert(13)

	n2, i2 := t.BTreeSearch(t.root, 13)
	fmt.Println("Key found:", "n:", n2, "i:", i2)

	t.BTreeInsert(7)
	fmt.Println("Non Empty B-tree:", t.root, t.root.children)
	t.BTreeInsert(24)
	fmt.Println("Non Empty B-tree:", t.root, t.root.children)
	fmt.Println("-=-=-=-=- post split:")
	t.BTreeInsert(16)
	fmt.Println("Non Empty B-tree:", t.root.children[0].keys, t.root.keys, t.root.children[1].keys)
	t.BTreeInsert(5)
	fmt.Println("Non Empty B-tree:", t.root.children[0].keys, t.root.keys, t.root.children[1].keys)
	t.BTreeInsert(50)
	t.BTreeInsert(51)
	fmt.Println("Non Empty B-tree:", t.root.children, t.root.children[0].keys, t.root.keys, t.root.children[1].keys, t.root.children[2].keys)
}
