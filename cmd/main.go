package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/hiimnhan/bible-quote/common"
	"github.com/hiimnhan/bible-quote/internal/indexing"
)

func main() {
	cm, err := common.ConstructCitationMap(common.KJVPath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	start := time.Now()
	index, err := indexing.InvertedIndex(cm)
	log.Printf("indexed finished in %v\n", time.Since(start))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	err = index.SaveToDisk(common.KJVInvertedIndexSavePath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	jsonFile, err := os.Open(common.KJVInvertedIndexSavePath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	bytesVal, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(bytesVal, &index)

	var query string
	flag.StringVar(&query, "q", "god", "search query")
	flag.Parse()

	start = time.Now()
	matchedIDs := index.Search(query)
	duration := time.Since(start)

	for _, id := range matchedIDs {
		fmt.Println(cm[id].String())

	}
	log.Printf("found %d matched in %v\n", len(matchedIDs), duration)

}
