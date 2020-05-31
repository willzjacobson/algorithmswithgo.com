package main

import "fmt"

type BinarySearchTreeNode struct {
	data  int
	p     *BinarySearchTreeNode
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func CreateBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// INSERTION
func (t *BinarySearchTree) BinarySearchTreeInsert(k int) {
	// Traverse tree to find who will become the new node's parent
	var p *BinarySearchTreeNode
	x := t.root
	for x != nil {
		p = x
		if k < x.data {
			x = x.left
		} else {
			x = x.right
		}
	}

	// create new node
	n := &BinarySearchTreeNode{
		data: k,
		p:    p,
	}

	// if tree is empty, new node is the root
	if p == nil {
		t.root = n
	} else if n.data < p.data {
		p.left = n
	} else {
		p.right = n
	}
}

// DELETION
func (t *BinarySearchTree) BinarySearchTreeDelete(k int) {
	n := t.BinarySearchTreeSearch(k)
	if n == nil {
		return
	}

	if n.left == nil {
		t.BinarySearchTreeTransplant(n, n.right)
	} else if n.right == nil {
		t.BinarySearchTreeTransplant(n, n.left)
	} else {
		// find the n's successor (we will swap it with n), which is the min of n's right subtree
		successor := t.BinarySearchTreeMinimum(n.right)
		if n.right != successor {
			// swap n's successor with its right child
			t.BinarySearchTreeTransplant(successor, n.right)
			successor.right = n.right
			successor.right.p = successor
		}
		// swap remove n by replacing it with it's successor, which we now know is its right child
		t.BinarySearchTreeTransplant(n, successor)
		successor.left = n.left
		n.left.p = successor
	}
}

// BinarySearchTreeTransplant : replaces a node as a child of its parent
func (t *BinarySearchTree) BinarySearchTreeTransplant(u, v *BinarySearchTreeNode) {
	if u.p == nil {
		t.root = v
	} else if u == u.p.right {
		u.p.right = v
	} else {
		u.p.left = v
	}

	if v != nil {
		v.p = u.p
	}
}

// SEARCHING
func (t *BinarySearchTree) BinarySearchTreeSearch(x int) *BinarySearchTreeNode {
	n := t.root
	for n != nil {
		if n.data == x {
			return n
		}
		if x < n.data {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}
func (t *BinarySearchTree) InOrderTreeWalk(n *BinarySearchTreeNode) {
	if n == nil {
		return
	}
	t.InOrderTreeWalk(n.left)
	fmt.Println(n.data)
	t.InOrderTreeWalk(n.right)
}
func (t *BinarySearchTree) PreOrderTreeWalk(n *BinarySearchTreeNode) {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	t.PreOrderTreeWalk(n.left)
	t.PreOrderTreeWalk(n.right)
}
func (t *BinarySearchTree) PostOrderTreeWalk(n *BinarySearchTreeNode) {
	if n == nil {
		return
	}
	t.PostOrderTreeWalk(n.left)
	t.PostOrderTreeWalk(n.right)
	fmt.Println(n.data)
}
func (t *BinarySearchTree) BinarySearchTreeBFS() {
	q := []*BinarySearchTreeNode{t.root}
	var next *BinarySearchTreeNode
	for len(q) > 0 {
		next = q[0]
		q = q[1:]

		if next == nil {
			continue
		}

		q = append(q, next.left, next.right)
	}
}
func (t *BinarySearchTree) BinarySearchTreeMinimum(n *BinarySearchTreeNode) *BinarySearchTreeNode {
	var p *BinarySearchTreeNode
	min := n
	for min != nil {
		p = min
		min = min.left
	}
	return p
}
func (t *BinarySearchTree) BinarySearchTreeMaximum(n *BinarySearchTreeNode) *BinarySearchTreeNode {
	var p *BinarySearchTreeNode
	max := n
	for max != nil {
		p = max
		max = max.right
	}
	return p
}

func (t *BinarySearchTree) BinarySearchTreePrecessor(n *BinarySearchTreeNode) *BinarySearchTreeNode {
	if n == nil {
		return n
	}
	// if node has a left child, it's precessor is the max of its left subtree
	if n.left != nil {
		return t.BinarySearchTreeMaximum(t.root.left)
	}

	// else, need to go up until we find the lowest ancestor whose right child is also an ancestor
	y := n.p
	for y != nil && n == y.left {
		n = y
		y = y.p
	}
	return y
}
func (t *BinarySearchTree) BinarySearchTreeSuccessor(n *BinarySearchTreeNode) *BinarySearchTreeNode {
	if n == nil {
		return n
	}

	// if node has a right child, it's successor is the min of its right subtree
	if n.right != nil {
		return t.BinarySearchTreeMinimum(n.right)
	}
	// else, need to go up until we find the lowest ancestor whose left child is also an ancestor
	y := n.p
	for y != nil && n == y.right {
		n = y
		y = y.p
	}
	return y
}

func main() {
	t := CreateBinarySearchTree()
	t.BinarySearchTreeInsert(5)
	t.BinarySearchTreeInsert(6)
	t.BinarySearchTreeInsert(4)
	t.BinarySearchTreeInsert(7)
	t.BinarySearchTreeInsert(3)
	t.BinarySearchTreeInsert(2)
	t.BinarySearchTreeInsert(4)
	t.BinarySearchTreeInsert(6)
	t.BinarySearchTreeInsert(8)
	t.InOrderTreeWalk(t.root)
	fmt.Println("-=-=-")
	t.PreOrderTreeWalk(t.root)
	fmt.Println("-=-=-")
	t.PostOrderTreeWalk(t.root)
	fmt.Println("-=-=-")
	t.BinarySearchTreeBFS()
	fmt.Println("-=-=-")
	fmt.Println("min:", t.BinarySearchTreeMinimum(t.root))
	fmt.Println("min:", t.BinarySearchTreeMinimum(t.root.left))
	fmt.Println("min:", t.BinarySearchTreeMinimum(t.root.right))
	fmt.Println("-=-=-")
	fmt.Println("max:", t.BinarySearchTreeMaximum(t.root))
	fmt.Println("-=-=-")
	fmt.Println("Precessor", t.BinarySearchTreePrecessor(t.root))
	fmt.Println("Precessor", t.BinarySearchTreePrecessor(t.root.right))
	fmt.Println("Precessor", t.BinarySearchTreePrecessor(t.root.right.right.left))
	fmt.Println("-=-=-")
	fmt.Println("Successor", t.BinarySearchTreeSuccessor(t.root))
	fmt.Println("Successor", t.BinarySearchTreeSuccessor(t.root.right))
	fmt.Println("Successor", t.BinarySearchTreeSuccessor(t.root.left))
	fmt.Println("Successor", t.BinarySearchTreeSuccessor(t.root.left.left))
	fmt.Println("-=-=-")
	t.BinarySearchTreeDelete(t.root.data)
	t.BinarySearchTreeDelete(t.root.left.left.left.data)
	t.BinarySearchTreeDelete(t.root.right.data)
	t.InOrderTreeWalk(t.root)
}
