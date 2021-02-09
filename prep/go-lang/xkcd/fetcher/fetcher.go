package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func Fetch() error {
	fmt.Printf("Hi from Fetch!\n\n")

	var req string

	f, err := os.Create("index.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	// Fetch a comic from xkcd site.
	// There are a total of 2422 comics!
	for i := 1; i <= 2; i++ {
		req = fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)

		resp, err := http.Get(req)
		if err != nil {
			return errors.Wrap(err, "failed to make GET request for xkcd comic")
		}

		// Read JSON payload from response.
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "failed to read body of response")
		}

		// Append two newline chars to the array to improve the
		// readability of data stored in file/index.
		data = append(data, 10)
		data = append(data, 10)

		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}
