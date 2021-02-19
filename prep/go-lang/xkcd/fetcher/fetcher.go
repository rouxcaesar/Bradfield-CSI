package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/pkg/errors"
)

// The current number of xkcd comics.
// TODO: Make this variable "discoverable" through
//       a network request to the xkcd website.
const maxComics = 2426

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

func Fetch() error {
	fmt.Printf("Hi from Fetch!\n\n")

	var req string
	var wg sync.WaitGroup
	messages := make(chan Comic, 10)
	done := make(chan bool)
	errs := make(chan error)
	wgDone := make(chan bool)
	index := make(map[int]string)

	for i := 1; i <= maxComics; i++ {
		//for i := 1; i <= 5; i++ {
		// We increment the wait group for each goroutine.
		wg.Add(1)

		// Spin off a separate producer goroutine to handle:
		//   - making a network request for a comic
		//   - parsing the response body
		//   - unmarshalling response body int a comic struct instance
		// The variable i is passed into the IIFE to avoid the common bug
		// involving variable scope and for loops.
		// Reference: https://dev.to/kkentzo/the-golang-for-loop-gotcha-1n35
		go func(i int) {
			fmt.Printf("In producer goroutine %d\n\n", i)

			// We decrement the wait group once the goroutine is finished.
			defer wg.Done()

			req = fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)

			resp, err := http.Get(req)
			if err != nil {
				// Need to send error into error channel
				fmt.Printf("err: %v\n", err)
				errs <- errors.Wrap(err, "failed to make GET request for xkcd comic")
				//return errors.Wrap(err, "failed to make GET request for xkcd comic")
			}

			// Read JSON payload from response.
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// Need to send error into error channel
				fmt.Printf("err: %v\n", err)
				errs <- errors.Wrap(err, "failed to read body of response")
				//return errors.Wrap(err, "failed to read body of response")
			}

			var comic Comic
			err = json.Unmarshal(data, &comic)

			// Pass constructed comic instance into channel for consumer
			// goroutine to process.
			messages <- comic
		}(i)
	}

	// Need to create a select stmt to listen to two channels:
	//  1) For completion of wait group (new channel and goroutine)
	//  2) For message from error channel
	// Block until all of the above goroutines finish.
	go func() {
		fmt.Printf("In wait group goroutine\n\n")
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case e := <-errs:
		close(errs)
		return e
	}

	// Close the channel to signal that we are done using it.
	// We can do this even before reading message from the channel.
	close(messages)
	fmt.Printf("Done with for loop\n\n")

	// This consumer goroutine reads messages off the channel and builds the
	// index map.
	// Maps are not safe for concurrency by themselves, so we couldn't have
	// the producer goroutines update the map.
	// With this approach, the map is updated in a sequential manner.
	go func() {
		fmt.Printf("Inside consumer goroutine\n\n")
		for c := range messages {
			index[c.Num] = c.Transcript
		}

		// Once we have finished reading all the messages on the channel,
		// we send a message on the done channel to indicate that we are
		// finished building the index.
		done <- true
	}()

	<-done

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
