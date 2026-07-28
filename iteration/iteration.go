package iteration

import "strings"

func Repeat(character string, amount int) string {
	var repeated strings.Builder
	for range amount {
		repeated.WriteString(character)
	}
	result := repeated.String()
	return result
}
