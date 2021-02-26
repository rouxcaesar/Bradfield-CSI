package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"bradfield-csi/prep/go-lang/xkcd/fetcher"
	"bradfield-csi/prep/go-lang/xkcd/index"
	"bradfield-csi/prep/go-lang/xkcd/search"
)

func main() {
	// Check if user wishes to view the help display.
	if os.Args[1] == "--help" {
		printHelp()
		os.Exit(0)
	}

	// Check args provided to program and grab search term if provided one.
	searchTerm, err := processArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check if offline index exists.
	// If not, we first build the offline index.
	if !index.IndexExists() {
		fmt.Printf("Offline index not found, building index now\n\n")

		// Add flag parsing to offer concurrent fetching?
		// --concurrent
		// Concurrent approach is giving unexpected results, going to
		// make this a stretch goal for the future.
		//err := fetcher.ConcurrentFetch()
		err := fetcher.Fetch()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		index.BuildIndex()
		fmt.Printf("Offline index built, ready to search\n\n")
	}

	// Load the offline index into memory for access by
	// opening the index.json file and decoding the
	// data into the index variable.
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

	// Argument searchTerm will be the argument passed in by the
	// user of this program.
	// Ex: `xkcd sheep` -> searchTerm == "sheep"
	err = search.SearchIndex(searchTerm, i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func printHelp() {
	fmt.Printf("xkcd - Search for xkcd comics with a provided search term\n\n")
	fmt.Printf("Usage: xkcd <string>\n\n")
	fmt.Printf("The search term can either be a number or a string of characters.\n")
	fmt.Printf("If it's a number, the program will search for the comics for one with the matching number based on the xkcd archive.\n")
	fmt.Printf("If it's a string of characters, the program will search for comics whose transcript contains the string.\n\n")
}

func processArgs() (string, error) {
	n := len(os.Args)

	if n < 1 {
		return "", errors.New("no search term provided, please see --help")
	} else if n > 2 {
		return "", errors.New("too many arguments provided, please see --help")
	}

	return os.Args[1], nil
}
