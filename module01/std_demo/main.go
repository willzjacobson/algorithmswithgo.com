package main

import (
	"algo/module01"
	"fmt"
)

/*
How to use:

Can run `go run std_demo/main.go` and enter an int,
then enter pairs of numbers manually

Can run `go run std_demo/main.go < std_demo/input.txt`
to feed in pairs of numbers from the input file

Can run `go run std_demo/main.go < std_demo/input.txt > std_demo/output.txt`
to feed in pairs of numbers from the input file and output them to a file
*/

func main() {
	// This int controls how many times the loop runs
	// (how many pairs of numbers this program is expecting)
	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module01.GCD(a, b)
		fmt.Println(gcd)
	}
}
