package module01

import (
	"strings"
)

var intToStr = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
}

// DecToBase will return a string representing
// the provided decimal number in the provided base.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	res := ""

	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}

	return res
}

// DecToBase2 will return a string representing
// the provided decimal number in the provided base.
// This implementation uses a string builder (more efficient),
// but has to reverse the string before returning it.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase2(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var sb strings.Builder

	for dec > 0 {
		rem := dec % base
		sb.WriteByte(charset[rem])
		// res = fmt.Sprintf("%X%s", rem, res) // this would be sufficient if we are capping the options for base at 16
		dec = dec / base
	}

	return Reverse(sb.String())
}

/*
Example:
num: 10
base: 2

10 % 2 = 0
"0"

10 / 2 = 5
5 % 2 = 1
"10"

5 / 2 = 2
2 % 2 = 0
"010"

2 / 2 = 1
1 % 2 = 1
"1010"

1 / 2 = 0
=> "1010"
*/
