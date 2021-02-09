package main

import (
	"fmt"
	"os"

	"bradfield-csi/prep/go-lang/xkcd/fetcher"
)

func main() {
	fmt.Printf("Hi from main!\n\n")
	err := fetcher.Fetch()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Finished execution, bye!\n")
}
