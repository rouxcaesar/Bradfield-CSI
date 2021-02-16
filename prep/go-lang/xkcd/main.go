package main

import (
	"encoding/json"
	"fmt"
	"os"

	"bradfield-csi/prep/go-lang/xkcd/fetcher"
	"bradfield-csi/prep/go-lang/xkcd/index"
	"bradfield-csi/prep/go-lang/xkcd/search"
)

// When printing the output of a found result, we need to
// show the URL and transcript.
// The transcript is part of the JSON response returned
// by fetching the URL, but the URL itself isn't in the
// response body payload.

func main() {
	fmt.Printf("Hi from main!\n\n")

	// Check args provided to program and grab search term.
	// searchTerm, err := processArgs()
	// if err != nil {
	//   fmt.Println(err)
	//   os.Exit(1)
	// }
	//
	// // Below check represents printing help statement.
	// // Could be better, don't want `nil` value to represent "--help".
	// if !searchTerm {
	//   os.Exit(0)
	// }

	// Check if offline index exists.
	// If not, output message to user and build index.
	// This will involve fetching the URLs, building the index,
	// and saving to a file.

	// 3) Next, parse argument to program which will be the search term.
	//    Search index for matching comics.
	//    For each match, write URL and transcript to output with newlines
	//    between each match..

	if !index.IndexExists() {

		fmt.Printf("Offline index not found, building index now\n")

		err := fetcher.Fetch()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		index.BuildIndex()

		fmt.Printf("Offline index built, ready to search\n")
	}

	// Now, load the offline index into memory for access.
	// Open index.json file and decode data into index variable.
	i := make(map[int]string)

	f, err := os.Open("index.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := json.NewDecoder(f).Decode(&i); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range i {
		fmt.Printf("%d: %s\n\n", k, v)
	}

	// Argument searchTerm will be the argument passed in by the
	// user of this program.
	//
	// Ex: `xkcd sheep` -> searchTerm == "sheep"
	//
	// To start, offer search based on comic number.
	// Ex: `xkcd 275` -> searchTerm == "275"
	searchTerm := "Baaaahhhhh"
	search.SearchIndex(searchTerm)

	fmt.Printf("Goodbye!\n")
}

func processArgs() (string, error) {
	if len(os.Args) > 2 {
		return nil, errors.New("too many arguments provided, please see --help")
	}

	if os.Args[1] == "--help" {
		printHelp()
		return nil, nil
	}
}
