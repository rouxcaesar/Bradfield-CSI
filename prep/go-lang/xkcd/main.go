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
//
// We can either:
// (1) Add the URL to the data stored in each value in the index.
// OR,
// (2) Have the URL as a constant and substitute in the `num` value
//     which is part of the response body.

func main() {
	fmt.Printf("Hi from main!\n\n")

	// 1) Check if offline index exists.
	//    If not, output message to user and build index.
	//			- This will involve fetching the URLs, building the index,
	//        and saving to a file.

	// 2) Now, load the index into memory for access.
	//    This will involve data serialization/deserialization.
	//    Take the data in the file and store into a map.

	// 3) Next, parse argument to program which will be the search term.
	//    Search index for matching comics.
	//    For each match, write URL and transcript to output with newlines
	//    between each match..

	if !index.IndexExists() {

		err := fetcher.Fetch()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		index.BuildIndex()

		fmt.Println("Back inside main!\n\n")

	}

	// Open index.json file and decode data into index variable.
	i := make(map[int]string)

	f, err := os.Open("index.json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if err := json.NewDecoder(f).Decode(&i); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for k, v := range i {
		fmt.Printf("%d: %s\n\n", k, v)
	}

	// Argument searchTerm will be the argument passed in by the
	// user of this program.
	//
	// Ex: `xkcd sheep` -> searchTerm == "sheep"
	searchTerm := "Baaaahhhhh"
	search.SearchIndex(searchTerm)

	fmt.Printf("Finished execution, bye!\n")
}
