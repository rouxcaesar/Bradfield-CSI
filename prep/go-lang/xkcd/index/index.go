package index

import (
	"fmt"
	"os"
)

// For our offline index of all the xkcd comics, we need
// a data structure that is suited for search operations
// on the data it contains.
//
// We don't really need to worry about the cost of insertions
// as we'll be fetching the comics only once.
//
// For the same reason, we don't need to worry about the time
// complexity for deletions or updates to the data strucure's contents.
//
// Reference: https://en.wikipedia.org/wiki/Search_data_structure
// For this program, we'll build a hash table for fast reads.

// BuildIndex will create the offline index.
func BuildIndex() {
	fmt.Printf("Inside BuildIndex!\n")
}

// IndexExists checks whether the offline index file
// already exists.
func IndexExists() bool {
	info, err := os.Stat("index.json")

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
