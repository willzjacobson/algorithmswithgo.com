package main

import (
	"fmt"
	"math"
)

func main() {
	str := "bb"
	res := longestPalindrome(str)
	fmt.Println("res:", res)

	str = "bab"
	res = longestPalindrome(str)
	fmt.Println("res:", res)

	str = "33baab"
	res = longestPalindrome(str)
	fmt.Println("res:", res)
}

// solved using the "expand around the center" technique
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start := 0
	end := 0
	for i := range s {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		l := int(math.Max(float64(l1), float64(l2)))
		if l > (end-start)+1 {
			if l%2 == 0 {
				start = i - l/2 + 1
			} else {
				start = i - l/2
			}
			end = i + l/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, start, end int) int {
	for start >= 0 && end < len(s) && string(s[start]) == string(s[end]) {
		if start >= 0 && end < len(s) {
			start--
			end++
		} else {
			break
		}
	}

	return end - start - 1
}
