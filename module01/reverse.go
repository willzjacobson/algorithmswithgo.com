package module01

import "strings"

// Reverse will return the provided word in reverse
// order by iterating through the string forwards. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(str string) string {
	newStr := ""
	for _, v := range str {
		newStr = string(v) + newStr
	}
	return newStr
}

// Reverse2 will return the provided word in reverse
// order by iterating through the string backwards. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse2(str string) string {
	var sb strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte((str[i]))
	}
	return sb.String()
}
