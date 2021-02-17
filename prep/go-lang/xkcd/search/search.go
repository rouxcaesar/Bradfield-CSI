package search

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	urlPrefix string = "https://xkcd.com/"
	urlSuffix string = "/info.0.json"
)

// SearchIndex takes a term provided by the user and
// searches the index for any entries (comics) whose
// comic number or transcript match the term.
func SearchIndex(searchTerm string, index map[int]string) error {
	//fmt.Printf("searchTerm: %s\n\n", searchTerm)

	num, err := strconv.Atoi(searchTerm)
	if err != nil {
		// Provided searchTerm is not a number.
		// Must search index for a string of characters.
		// searchTranscripts(searchTerm, index)
		return errors.New("search based on string of chars not implemented yet")
	}

	return searchKeys(num, index)
}

func searchKeys(num int, index map[int]string) error {
	v, ok := index[num]
	if !ok {
		return errors.New("No comic found for that number")
	}

	printMatch(num, v)
	return nil
}

// printMatch will take a match from the search and print it
// to stdout for the user to view.
func printMatch(indexKey int, transcript string) {
	fullURL := fmt.Sprintf(urlPrefix+"%d"+urlSuffix, indexKey)
	fmt.Printf("\n" + fullURL + "\n\n" + transcript + "\n\n")
}
