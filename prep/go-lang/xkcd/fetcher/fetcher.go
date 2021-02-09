package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func Fetch() error {
	fmt.Printf("Hi from Fetch!\n\n")

	// Fetch a comic from xkcd site.
	resp, err := http.Get("https://xkcd.com/571/info.0.json")
	if err != nil {
		return errors.Wrap(err, "failed to make GET request for xkcd comic")
	}

	// Read JSON payload from response.
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read body of response")
	}

	fmt.Printf("%s\n\n", data)
	return nil
}
