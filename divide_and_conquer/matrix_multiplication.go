package main

import "fmt"

func main() {
	m1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	m2 := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	}

	res := multiplyMatricesBrute(m1, m2)
	fmt.Println("res:", res)
}

// handles nxm matrix; does not have to be nxn
func multiplyMatricesBrute(m1 [][]int, m2 [][]int) [][]int {
	l := len(m1)
	if l != len(m2) {
		panic("Cannot multiply these 2 matrices")
	}

	m := [][]int{}
	for i := 0; i < l; i++ {
		m = append(m, []int{})
	}

	for y := range m1 {
		row := m1[y]
		for x := range m2[0] {
			newYX := 0
			for i, v := range row {
				newYX += v * m2[i][x]
			}
			m[y] = append(m[y], newYX)
		}
	}
	return m
}

// Strassen's method
// Runs in n^lg7 (or n^~2.81) time, which is
// asymptitically better than the Square Matrix Multiply (shown above)
// K, fuck that, I don't want to code it out. But see p. 100 of Intro to Algorithms course
