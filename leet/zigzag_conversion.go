package main

import (
	"fmt"
	"math"
)

/*
P   A   H   N
A P L S I I G
Y   I   R

==> "PAHNAPLSIIGYIR"
*/

func main() {
	s := "PAYPALISHIRING"
	res := convert(s, 3)
	fmt.Println("res:", res)

	res2 := convert(s, 4)
	fmt.Println("res2:", res2)

	res3 := convert("AB", 1)
	fmt.Println("res3:", res3)
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	res := ""
	rows := int(math.Min(float64(numRows), float64(len(s))))
	m := []string{}
	for i := 0; i < rows; i++ {
		m = append(m, "")
	}

	curRow := 0
	goingDown := false
	for _, r := range s {
		// Add string to appropriate slice in the matrix
		m[curRow] += string(r)

		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}

	}

	for _, v := range m {
		res += v
	}

	return res
}
