package iteration

import "strings"

func Repeat(s string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}

func Matches(child, parent string) bool {
	return strings.Contains(parent, child)
}
