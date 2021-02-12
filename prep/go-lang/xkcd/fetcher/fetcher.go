package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

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
	index := make(map[int]string)

	//f, err := os.Create("index.txt")
	//if err != nil {
	//	return err
	//}
	//defer f.Close()

	// Fetch a comic from xkcd site.
	// There are a total of 2422 comics!
	//
	// TODO: Make the fetching of URLs concurrent with
	//       goroutines and send data through channel.
	//       Then on other end of channel read data and
	//       write data to file/index.
	//for i := 1; i <= 2422; i++ {
	for i := 1; i <= 5; i++ {
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

		var comic Comic
		err = json.Unmarshal(data, &comic)
		//fmt.Printf("%+v\n\n", comic)

		index[comic.Num] = comic.Transcript

		// Append two newline chars to the array to improve the
		// readability of data stored in file/index.
		//data = append(data, 10)
		//data = append(data, 10)

		//	_, err = f.Write(data)
		//	if err != nil {
		//		return err
		//	}

		//fmt.Printf("%s\n\n", data)
	}

	//	for k, v := range index {
	//		fmt.Printf("%d: %s\n\n", k, v)
	//	}

	// Now save the contents of the index variable to a file
	// to make it an "offline" index.
	file, err := os.Create("index.json")
	if err != nil {
		return errors.Wrap(err, "failed to create file index.json")
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(&index); err != nil {
		return errors.Wrap(err, "failed to encode index into index.json")
	}

	return nil
}
