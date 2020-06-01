package main

import "fmt"

const (
	red   = "red"
	black = "black"
)

type RBTreeNode struct {
	data  int
	color string
	p     *RBTreeNode
	left  *RBTreeNode
	right *RBTreeNode
}

type RBTree struct {
	root *RBTreeNode
	null *RBTreeNode
}

func CreateRBTree() *RBTree {
	null := &RBTreeNode{}
	return &RBTree{
		root: null,
		null: null,
	}
}

// INSERTION
func (t *RBTree) RBTreeInsert(k int) {
	// Traverse tree to find who will become the new node's parent
	var p = t.null
	x := t.root
	for x != t.null {
		p = x
		if k < x.data {
			x = x.left
		} else {
			x = x.right
		}
	}

	// create new node
	n := &RBTreeNode{
		data:  k,
		p:     p,
		left:  t.null,
		right: t.null,
		color: red,
	}

	// if tree is empty, new node is the root
	if p == t.null {
		t.root = n
	} else if n.data < p.data {
		p.left = n
	} else {
		p.right = n
	}

	// Ensure red-black rules are upheld
	t.InsertFixup(n)
}

func (t *RBTree) InsertFixup(n *RBTreeNode) {
	for n.p.color == red {
		if n.p == n.p.p.left {
			uncle := n.p.p.right
			if uncle.color == red {
				n.p.color = black
				uncle.color = black
				n.p.p.color = red
				n = n.p.p
			} else {
				if n == n.p.right {
					n = n.p
					t.LeftRotate(n)
				}
				n.p.color = black
				n.p.p.color = red
				t.RightRotate(n.p.p)
			}
		} else {
			// mirror image of the above case
			uncle := n.p.p.left
			if uncle.color == red {
				n.p.color = black
				uncle.color = black
				n.p.p.color = red
				n = n.p.p
			} else {
				if n == n.p.left {
					n = n.p
					t.RightRotate(n)
				}

				n.p.color = black
				n.p.p.color = red
				t.LeftRotate(n.p.p)
			}
		}
	}
	t.root.color = black
}

// DELETION (untested, but I believe correct)
// Deletion is complex. I'm not likely to rememember how it works
func (t *RBTree) RBTreeDelete(k int) {
	n := t.RBTreeSearch(k)
	if n == t.null {
		return
	}

	y := n // y will take the place of n
	yOriginalColor := y.color
	var x *RBTreeNode // when y takes the place of n, x will take the place of y

	if n.left == t.null {
		x = n.right
		t.RBTreeTransplant(n, n.right)
	} else if n.right == t.null {
		x = n.left
		t.RBTreeTransplant(n, n.left)
	} else {
		// if n has 2 children, we will replace n with its successor; make y n's successor
		y = t.RBTreeMinimum(n.right)
		yOriginalColor = y.color
		x = y.right
		// If y is n's right child, y will be x's parent, since x is moving to y's old spot
		if y.p == n {
			x.p = y
		} else {
			t.RBTreeTransplant(y, y.right)
			y.right = n.right
			y.right.p = y
		}

		// swap remove n by replacing it with it's successor, which we now know is its right child
		t.RBTreeTransplant(n, y)
		y.left = n.left
		n.left.p = y
		y.color = n.color
	}

	if yOriginalColor == black {
		t.DeleteFixup(x)
	}
}

// RBTreeTransplant : replaces a node as a child of its parent
func (t *RBTree) RBTreeTransplant(u, v *RBTreeNode) {
	if u.p == t.null {
		t.root = v
	} else if u == u.p.right {
		u.p.right = v
	} else {
		u.p.left = v
	}

	v.p = u.p
}

func (t *RBTree) DeleteFixup(n *RBTreeNode) {
	for n != t.root && n.color == black {
		if n == n.p.left {
			w := n.p.right // n's sibling
			// case 1 (can then become case 2)
			if n.color == red {
				w.color = black
				n.p.color = red
				t.LeftRotate(n.p)
				w = n.p.right
			}
			// case 2
			if w.left.color == black && w.right.color == black {
				w.color = red
				n = n.p
				// case 3 (is transformed into case 4)
			} else {
				if w.right.color == black {
					w.left.color = black
					w.right.color = black
					t.RightRotate(w)
					w = n.p.right
				}
				// case 4
				w.color = n.p.color
				n.p.color = black
				w.right.color = black
				t.LeftRotate(n.p)
				n = t.root // to terminate the loop
			}
		} else {
			w := n.p.left // n's sibling
			// case 1 (can then become case 2)
			if n.color == red {
				w.color = black
				n.p.color = red
				t.RightRotate(n.p)
				w = n.p.left
			}
			// case 2
			if w.left.color == black && w.right.color == black {
				w.color = red
				n = n.p
			} else {
				// case 3 (is transformed into case 4)
				if w.left.color == black {
					w.right.color = black
					w.left.color = black
					t.LeftRotate(w)
					w = n.p.left
				}
				// case 4
				w.color = n.p.color
				n.p.color = black
				w.left.color = black
				t.RightRotate(n.p)
				n = t.root // to terminate the loop
			}
		}
	}
	n.color = black
}

// ROTATION
// LeftRotate : perform left rotation about a node on the tree
func (t *RBTree) LeftRotate(x *RBTreeNode) {
	y := x.right
	if y == t.null {
		return
	}

	y.p = x.p
	x.right = y.left
	if y.left != t.null {
		y.left.p = x
	}
	y.left = x

	if x.p == t.null {
		t.root = y
	} else if x.p.left == x {
		x.p.left = y
	} else {
		x.p.right = y
	}
	x.p = y
}

// RightRotate : perform right rotation about a node on the tree
func (t *RBTree) RightRotate(x *RBTreeNode) {
	y := x.left
	if y == t.null {
		return
	}

	y.p = x.p
	x.left = y.right
	if y.right != t.null {
		y.right.p = x
	}
	y.right = x

	if x.p == t.null {
		t.root = y
	} else if x.p.left == x {
		x.p.left = y
	} else {
		x.p.right = y
	}
	x.p = y
}

// SEARCHING (Same as for BinarySearchTree)
func (t *RBTree) RBTreeSearch(x int) *RBTreeNode {
	n := t.root
	for n != t.null {
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
func (t *RBTree) RBInOrderTreeWalk(n *RBTreeNode) {
	if n == t.null {
		return
	}
	t.RBInOrderTreeWalk(n.left)
	fmt.Println(n.data)
	t.RBInOrderTreeWalk(n.right)
}
func (t *RBTree) RBTreeMinimum(n *RBTreeNode) *RBTreeNode {
	var p *RBTreeNode
	min := n
	for min != t.null {
		p = min
		min = min.left
	}
	return p
}

func main() {
	t := CreateRBTree()
	t.RBTreeInsert(11)
	t.RBTreeInsert(2)
	t.RBTreeInsert(1)
	t.RBTreeInsert(7)
	t.RBTreeInsert(5)
	t.RBTreeInsert(4)
	t.RBTreeInsert(8)
	t.RBTreeInsert(14)
	t.RBTreeInsert(15)

	fmt.Println("Root", t.root)
	fmt.Println("Root.right", t.root.right)
	fmt.Println("Root.right.right", t.root.right.right)
	fmt.Println("Root.right.right.right", t.root.right.right.right)
	fmt.Println("Root.right.right.left.right", t.root.right.right.left.right)
	fmt.Println("Root.right.left", t.root.right.left)
	fmt.Println("Root.right.left.left", t.root.right.left.left)
	fmt.Println("-=-=-")
	fmt.Println("Root", t.root)
	fmt.Println("Root.left", t.root.left)
	fmt.Println("Root.left.left", t.root.left.left)

}
