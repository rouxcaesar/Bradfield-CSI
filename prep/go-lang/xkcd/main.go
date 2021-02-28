package main

import (
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
		handleError(err)
	}

	// Check if offline index exists.
	// If not, we first build the offline index.
	if !index.Exists() {
		fmt.Printf("Offline index not found, building index now\n\n")

		comics, err := fetcher.Fetch()
		if err != nil {
			handleError(err)
		}

		err = index.Build(comics)
		if err != nil {
			handleError(err)
		}

		fmt.Printf("Offline index built, ready to search\n\n")
	}

	offlineIndex, err := index.Load()
	if err != nil {
		handleError(err)
	}

	// Argument searchTerm will be the argument passed in by the
	// user of this program.
	// Ex: `xkcd sheep` -> searchTerm == "sheep"
	err = search.Index(searchTerm, offlineIndex)
	if err != nil {
		handleError(err)
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

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
