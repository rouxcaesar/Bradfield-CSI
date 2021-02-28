package index

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
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

// Build creates an offline index by writing the contents
// of a map into a local file.
func Build(data map[int]string) error {
	fmt.Printf("Inside BuildIndex!\n")
	// Now save the contents of the index variable to a file
	// to make it an "offline" index.
	file, err := os.Create(indexName)
	if err != nil {
		return errors.Wrap(err, "failed to create offline index")
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(&data); err != nil {
		return errors.Wrap(err, "failed to encode data into offline index")
	}

	return nil
}

// Exists checks whether the offline index file
// already exists.
func Exists() bool {
	info, err := os.Stat(indexName)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// Load opens the offline index file and decodes the contents
// into an in-memory map for use.
func Load() (map[int]string, error) {
	i := make(map[int]string)

	f, err := os.Open(indexName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file containing offline index")
	}

	if err := json.NewDecoder(f).Decode(&i); err != nil {
		return nil, errors.Wrap(err, "failed to decode file containing offline index")
	}

	return i, nil
}
