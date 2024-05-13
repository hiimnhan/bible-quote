package indexing

import (
	"log"

	"github.com/hiimnhan/bible-quote/common"
	"github.com/kylelemons/godebug/pretty"
	// "github.com/kylelemons/godebug/pretty"
)

type Set map[string]bool
type Index map[string]Set

func InvertedIndex(cm common.CitationMap) (*Index, error) {
	log.Println("Indexing...")
	idx := make(Index)

	for _, v := range cm {
		for _, token := range common.TokenizeAndFilter(v.Text) {
			set := idx[token]
			if set == nil {
				set = Set{}
			}
			if set[v.ID] {
				continue
			}
			set[v.ID] = true
			idx[token] = set
			pretty.Print(token, set)
		}
	}

	return &idx, nil
}

func intersection(s1, s2 []string) (is []string) {
	log.Printf("s1 %s\n", s1)
	log.Printf("s2 %s\n", s2)
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}

	for _, e := range s2 {
		if hash[e] {
			is = append(is, e)
		}
	}

	return
}

func (idx Index) Search(text string) []string {
	log.Printf("querying %s\n", text)
	var res []string
	for _, token := range common.TokenizeAndFilter(text) {
		log.Printf("token %s\n", token)
		if ids, ok := idx[token]; ok {
			var ss []string
			for s := range ids {
				ss = append(ss, s)
			}
			if res == nil {
				res = ss
			} else {
				res = intersection(res, ss)
			}
		} else {
			return nil
		}
	}

	return res
}
