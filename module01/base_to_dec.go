package module01

var strToInt = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
// currently only works up to base 16 (limited by the hardcoded charset)
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	res := 0
	multiplier := 1 // base^0

	for i := len(value) - 1; i >= 0; i-- {
		factor := strToInt[string(value[i])]
		res += factor * multiplier
		multiplier *= base
	}

	return int(res)
}

/*
the number 123 in base 10
1     2     3
10^2  10^1  10^0
1*10^2 + 2*10^1 + 3*10^0 = 123

the number 14 in base 2
1     1     1     0
2^3   2^2   2^1   2^0
1*2^3 + 1*2^2 + 1*2^1 + 0*2^0 = 14
*/
