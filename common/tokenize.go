package common

import (
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func TokenizeAndFilter(text string) []string {
	tokens := tokenize(text)
	tokens = lowercase(tokens)
	// tokens = stopwordFilter(tokens)
	return tokens
}
