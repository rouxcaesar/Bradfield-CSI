package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// The current number of xkcd comics.
// TODO: Make this variable "discoverable" through
//       a network request to the xkcd website.
const (
	maxComics = 2429
	// If concurrencyLimit is greater, we have weird situations in which not all the comics
	// are properly retrieved and stored - searches of index return fewer matches than there should be.
	concurrencyLimit = 10
)

// Comic is used to store the decoded JSON data representing a comic as returned in the response
// from a request to the xkcd website made by Fetch().
type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// ConcurrentFetch is a WIP version of Fetch() that uses concurrency to make concurrent requests
// to the xkcd website for comic data.
// The goal is to speed up the slowest portion of this program which is making all the network requests
// and handling the responses.
// This is a strech goal, I'll circle back to flesh this out in a later point in time.
func ConcurrentFetch() error {
	fmt.Printf("Hi from ConcurrentFetch!\n\n")

	var req string
	index := make(map[int]string)
	//	dataChan := make(chan []byte, 20)
	comicChan := make(chan Comic, 20)
	var wg sync.WaitGroup

	fmt.Printf("Starting goroutines\n\n")

	for i := 1; i <= 10; i++ {
		if i == 404 {
			continue
		}

		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			req = fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
			fmt.Printf("Req: %s\n", req)

			resp, err := http.Get(req)
			if err != nil {
				log.Printf("failed to make GET request - err: %v\n", err)
				log.Printf("URL: %s\n\n", req)
				//return errors.Wrap(err, "failed to make GET request for xkcd comic")
			}

			//if resp.StatusCode != http.StatusOK {
			//	resp.Body.Close()
			//	log.Printf("GET request returned non-OK status code for comic %d", num)
			//}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("failed to read response body - err: %v\n", err)
			}
			resp.Body.Close()

			var comic Comic
			err = json.Unmarshal(data, &comic)
			if err != nil {
				log.Printf("error decoding response: %v", err)
				if e, ok := err.(*json.SyntaxError); ok {
					log.Printf("syntax error at byte offset %d", e.Offset)
				}
				log.Printf("response: %q", data)
				log.Println("failed to unmarshal JSON")
				//return errors.Wrap(err, "failed to unmarshal JSON")
			}

			comicChan <- comic
		}(i)
	}

	wg.Wait()
	close(comicChan)

	for c := range comicChan {
		fmt.Printf("comic %d\n%s\n\n", c.Num, c.Transcript)
		index[c.Num] = c.Transcript
	}

	fmt.Printf("About to create file\n\n")

	// Now save the contents of the index variable to a file
	// to make it an "offline" index.
	// This should be moved to the BuildIndex() func in the index package.
	file, err := os.Create("index.json")
	if err != nil {
		return errors.Wrap(err, "failed to create file index.json")
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(&index); err != nil {
		return errors.Wrap(err, "failed to encode index into index.json")
	}

	fmt.Printf("Done with Fetcher func!\n\n")
	return nil

}

// Fetch makes requests to the xkcd website for all existing comics.
// It will take the JSON response for each and store the needed data
// into a map, which will be stored in an offline index using the
// index package.
func Fetch() (map[int]string, error) {
	fmt.Printf("Hi from Fetch!\n\n")

	go spinner(100 * time.Millisecond)

	var req string
	comics := make(map[int]string)

	for i := 1; i <= maxComics; i++ {
		// There is no comic number 404, it simply returns a 404: Not Found response.
		// This could be the underlying issue causing my concurrent approach to fail!
		if i == 404 {
			continue
		}

		req = fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		resp, err := http.Get(req)
		if err != nil {
			log.Printf("err: %v\n", err)
			log.Printf("URL: %s\n\n", req)
			return nil, errors.Wrap(err, "failed to make GET request for xkcd comic")
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			log.Printf("GET request returned non-OK status code for comic %d", i)
		}

		var comic Comic

		// Read JSON payload from response.
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("err: %v\n", err)
			log.Printf("data: %v\n\n", data)
			return nil, errors.Wrap(err, "failed to read body of response")
		}
		resp.Body.Close()

		err = json.Unmarshal(data, &comic)
		if err != nil {
			log.Printf("i: %d", i)
			log.Printf("error decoding response: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("response: %q", data)
			return nil, errors.Wrap(err, "failed to unmarshal JSON")
		}

		// Sometimes the transcript value in the JSON response is an empty string.
		// In these cases, use the alt value in the response instead.
		if comic.Transcript != "" {
			comics[comic.Num] = comic.Transcript
		} else {
			comics[comic.Num] = comic.Alt
		}
	}

	fmt.Printf("Done with Fetcher func!\n\n")
	return comics, nil
}

// spinner outputs a visual indicator that the program
// is currently fetching comics.
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
