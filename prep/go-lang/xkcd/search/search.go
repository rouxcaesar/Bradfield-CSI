package search

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	urlPrefix string = "https://xkcd.com/"
	urlSuffix string = "/info.0.json"
)

// Index takes a term provided by the user and
// searches the index for any entries (comics) whose
// comic number or transcript match the term.
func Index(searchTerm string, index map[int]string) error {
	//fmt.Printf("searchTerm: %s\n\n", searchTerm)

	num, err := strconv.Atoi(searchTerm)
	if err != nil {
		// Provided searchTerm is not a number.
		// Must search index for a string of characters.
		return searchTranscripts(searchTerm, index)
	}

	return searchKeys(num, index)
}

// searchTranscripts will consider each value stored in the index
// and search for a match with the user provided search term.
// This search term will be a string of characters.
func searchTranscripts(searchTerm string, index map[int]string) error {
	found := false

	// Search every comic in the index for a match in with provided searchTerm.
	for k, v := range index {
		if strings.Contains(v, searchTerm) {
			printMatch(k, v)
			found = true
		}
	}

	if !found {
		return errors.New("no comic was found for provided search term")
	}

	return nil
}

// searchKeys will look up the transcript of a comic
// that is stored in the offline index via the comic's number.
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
