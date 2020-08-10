package main

import (
	"fmt"
)

func buildBadMatchTable(pattern string) map[string]int {
	badMatchTable := map[string]int{}
	patternLen := len(pattern)

	for i, v := range pattern {
		// how many characters we should shift if this character is encountered
		badMatchTable[string(v)] = patternLen - 1 - i
	}
	fmt.Println("Bad match table: ", badMatchTable)
	return badMatchTable
}

// BMH : Looking for "pattern" within "toSearch"
// Performance (n is length of string to search, m is length of pattern to match)
//   best case: O(n/m)
//   worst case: O(nm)
func BMH(toSearch, pattern string) bool {
	toSearchLen := len(toSearch)
	patternLen := len(pattern)
	badMatchTable := buildBadMatchTable(pattern)

	// start evaluating at the last character index of pattern
	currentInd := patternLen - 1

	for currentInd < toSearchLen {
		numCharsMatched := 0

		for string(pattern[patternLen-1-numCharsMatched]) == string(toSearch[currentInd-numCharsMatched]) {
			numCharsMatched++
			if numCharsMatched == patternLen {
				return true
			}
		}
		val := badMatchTable[string(toSearch[currentInd])]
		if val == 0 {
			currentInd += patternLen
		}
		currentInd += val
	}
	return false
}

func main() {
	toSearch := "We hold these truth" // s to be self-evident"
	pattern := "truth"

	fmt.Println("haveMatch:", BMH(toSearch, pattern))
}
