package indexing

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hiimnhan/bible-quote/common"
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
				set = make(Set)
			}
			if set[v.ID] {
				continue
			}
			set[v.ID] = true
			idx[token] = set
		}
	}

	return &idx, nil
}

func (index *Index) SaveToDisk(path string) error {
	log.Printf("Saving to %s...\n", path)
	file, err := json.MarshalIndent(&index, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, file, 0666)

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

func (idx *Index) Search(text string) []string {
	log.Printf("querying %s\n", text)
	var res []string
	for _, token := range common.TokenizeAndFilter(text) {
		log.Printf("token %s\n", token)
		if ids, ok := (*idx)[token]; ok {
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
