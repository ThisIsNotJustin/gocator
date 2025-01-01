package utils

import (
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	isDelimiter := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}

	return strings.FieldsFunc(text, isDelimiter)
}

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseTokenSlice(tokens)
	tokens = removeFilterSlice(tokens)
	tokens = stemmedFilterSlice(tokens)
	return tokens
}
