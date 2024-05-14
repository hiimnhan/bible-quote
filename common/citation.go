package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "github.com/google/uuid"
	// "github.com/kylelemons/godebug/pretty"
)

type Citation struct {
	Book    string
	Chapter int
	Verse   int
	Text    string
	ID      string
}

type CitationMap map[string]*Citation

func generateCitations(path string) ([]Citation, error) {
	f, err := os.Open(path)
	if err != nil {
		return []Citation{}, err
	}
	defer f.Close()

	var citations []Citation

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		c, err := serializeCitation(line)
		if err != nil {
			return []Citation{}, err
		}
		citations = append(citations, c)
	}

	return citations, nil
}

func ConstructCitationMap(path string) (CitationMap, error) {
	citations, err := generateCitations(path)
	if err != nil {
		return CitationMap{}, err
	}

	citationMap := make(CitationMap)
	for _, c := range citations {
		citationMap[c.ID] = &c
	}

	return citationMap, nil
}

func serializeCitation(line string) (Citation, error) {
	parts := strings.SplitN(strings.ReplaceAll(line, "\t", " "), " ", 3)
	if len(parts) < 3 {
		return Citation{}, errors.New("Wrong format line")
	}

	chapterAndVerse := strings.Split(parts[1], ":")
	if len(chapterAndVerse) < 2 {
		return Citation{}, errors.New("Wrong format line")
	}

	chapter, err := strconv.Atoi(chapterAndVerse[0])
	if err != nil {
		return Citation{}, errors.New("Wrong format chapter")
	}

	verse, err := strconv.Atoi(chapterAndVerse[1])
	if err != nil {
		return Citation{}, errors.New("Wrong format verse")
	}

	return Citation{
		Book:    parts[0],
		Chapter: chapter,
		Verse:   verse,
		Text:    parts[2],
		// ID:      uuid.New().String(),
		ID: strings.Join([]string{parts[0], fmt.Sprintf("%d", chapter), fmt.Sprintf("%d", verse)}, "-"),
	}, nil
}

func (c *Citation) String() string {
	return fmt.Sprintf("King James Bible, %s %d:%d\n%s\n\n", c.Book, c.Chapter, c.Verse, c.Text)
}
