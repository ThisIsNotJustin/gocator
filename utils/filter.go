package utils

import (
	"strings"
	"sync"

	snowballeng "github.com/kljensen/snowball/english"
)

var stopWords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	"or": {}, "is": {},
}

func lowercaseTokenSlice(tokens []string) []string {
	tokenSlice := make([]string, len(tokens))
	for i, token := range tokens {
		tokenSlice[i] = strings.ToLower(token)
	}

	return tokenSlice
}

func removeFilterSlice(tokens []string) []string {
	filtered := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, exists := stopWords[token]; !exists {
			filtered = append(filtered, token)
		}
	}

	return filtered
}

func stemmedFilterSlice(tokens []string) []string {
	tokenSlice := make([]string, len(tokens))
	var waitGroup sync.WaitGroup
	for i, token := range tokens {
		waitGroup.Add(1)
		go func(i int, token string) {
			defer waitGroup.Done()
			tokenSlice[i] = snowballeng.Stem(token, false)
		}(i, token)

	}

	waitGroup.Wait()
	return tokenSlice
}
