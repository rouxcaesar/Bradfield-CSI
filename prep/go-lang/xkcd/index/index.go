package index

import (
	"fmt"
	"os"
)

const indexName = "offline-index.json"

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
// Thus, a hash map will provide us with a simple data structure
// that features fast read and insert operations.

// BuildIndex will create the offline index.
// Right now, this logic is coupled within fetch.go
func BuildIndex() {
	fmt.Printf("Inside BuildIndex!\n")
}

// IndexExists checks whether the offline index file
// already exists.
func IndexExists() bool {
	info, err := os.Stat(indexName)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
