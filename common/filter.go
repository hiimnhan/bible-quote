package common

import "strings"

func lowercase(tokens []string) []string {
	lc := make([]string, len(tokens))
	for i, token := range tokens {
		lc[i] = strings.ToLower(token)
	}

	return lc
}

// filter top 10 commonly used words
// var stopwords = map[string]struct{}{
// 	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
// 	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
// }
//
// func stopwordFilter(tokens []string) []string {
// 	r := make([]string, 0, len(tokens))
// 	for _, token := range tokens {
// 		if _, ok := stopwords[token]; !ok {
// 			r = append(r, token)
// 		}
// 	}
// 	return r
// }
