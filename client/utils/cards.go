package utils

import "strings"

func SeparateByFour(s string) string {
	s = strings.ReplaceAll(s, "-", "")
	var builder strings.Builder
	length := len(s)
	for i := 0; i < length; i += 4 {
		if i > 0 {
			builder.WriteString("-")
		}
		if i+4 < length {
			builder.WriteString(s[i : i+4])
		} else {
			builder.WriteString(s[i:])
		}
	}
	return builder.String()
}
