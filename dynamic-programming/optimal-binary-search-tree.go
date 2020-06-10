package main

import "fmt"

/*
The Optimal Binary Search Tree problem
Of all the nodes in the tree, each node k has a certain probability that it is being looked up. Similarly, we may be looking for a node that is not present. We represent this by having a node d between each node k, such that di represents a search for all values between ki-1 and ki. Each hypothetical node d also has a probability of being searched for.
Given a list of nodes k(1-n) and d(0-n), we need to come up with the optimal binary search tree, with the minimum expected cost.
The expected cost of a search is the sum of the (depth of each node +1) * the probability of searching for that node.
See P.401 of Intro To Algorithms to see how to represent this mathamatically.
*/

// OptimalBST : Determine the lowest cost binary search tree, given a list of nodes and the probabilities of needing to look up
func OptimalBST(p, q []float32) ([][]float32, [][]int) {
	n := len(p)
	c := make([][]float32, n+1) // table to hold the minimum cost for each subtree
	w := make([][]float32, n+1) // table to hold the sum of the probabilities of all the nodes in a subtree (used in calculations fo table c)
	root := make([][]int, n)    // table to hold root node resulting in the minimum cost for subtree
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			c[i] = append(c[i], 0)
			w[i] = append(w[i], 0)
			if i < n && j < n {
				root[i] = append(root[i], -1)
			}
		}
	}

	for i := 0; i <= n; i++ {
		c[i][i] = q[i]
		w[i][i] = q[i]
	}

	for l := 1; l <= n; l++ { // subtrees of different sizes
		for i := 1; i <= n-l+1; i++ { // nodes in the subtree in this subproblem starts at p[i]
			j := i + l - 1                          // final node p[j] that is included in the subtree of size l starting at index p[i]
			w[i-1][j] = w[i-1][j-1] + p[j-1] + q[j] // calculating sum of probabilities of nodes using that of this subtree minus this node, plus the probability of node p[j] and all nodes between p[j] and p[j+1]

			for r := i; r <= j; r++ { // trying out different roots for this subtree to see which yields optimal cost
				q := c[i-1][r-1] + c[r][j] + w[i-1][j] // calculate expected cost for subtree if we make p[r] the root

				// if cost is best so far, use p[r] as the root for this subtree
				if c[i-1][j] == 0 || q < c[i-1][j] {
					c[i-1][j] = q
					root[i-1][j-1] = r
				}
			}

		}
	}

	return c, root
}

// ConstructOptimalBST : Given the optomal root nodes for the subtrees, prints out the tree node by node.
// UNFINISHED
func ConstructOptimalBST(root [][]int, i, j int) {
	r := root[i][j]
	fmt.Println(r)

	leftStart := i
	leftEnd := r - 1 // not -1 cuz accounting for 0 indexing
	if leftStart < leftEnd {
		ConstructOptimalBST(root, leftStart, leftEnd)
		fmt.Println("going again left", leftStart, leftEnd)
	}

	rightStart := r // not +1 cuz accounting for 0 indexing
	rightEnd := j
	if rightStart < rightEnd {
		// ConstructOptimalBST(root, rightStart, rightEnd)
		fmt.Println("going again right", rightStart, rightEnd)
	}
}

func main() {
	p := []float32{.15, .1, .05, .1, .2}
	q := []float32{.05, .1, .05, .05, .05, .1}

	c, root := OptimalBST(p, q)
	fmt.Println("-=-=-")
	fmt.Println("cost matrix:", c)
	fmt.Println("root:", root)
	fmt.Println("-=-=-")
	ConstructOptimalBST(root, 0, len(p)-1)
}
