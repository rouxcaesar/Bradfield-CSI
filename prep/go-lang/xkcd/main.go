package main

import (
	"fmt"

	"bradfield-csi/prep/go-lang/xkcd/fetcher"
)

func main() {
	fmt.Printf("Hi from CSI!\n")
	fetcher.Fetch()
}
