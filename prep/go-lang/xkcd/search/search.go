package search

import "fmt"

const (
	urlPrefix string = "https://xkcd.com/"
	urlSuffix string = "/info.0.json"
)

// SearchIndex takes a term provided by the user and
// searches the index for any entries (comics) whose
// comic number or transcript match the term.
func SearchIndex(searchTerm string) {
	fmt.Printf("searchTerm: %s\n\n", searchTerm)
}

// printMatch will take a match from the search and print it
// to stdout for the user to view.
func printMatch(indexKey int, index map[int]string) {
	fullURL := fmt.Sprintf(urlPrefix+"%d"+urlSuffix, indexKey)

	fmt.Println(fullURL + "\n")
	fmt.Println(index[indexKey] + "\n")
}
