# xkcd

A program that fetches all of the `xkcd` comics, builds an offline index, and allows the user to search the index for a comic.

## Benchmarks

Data collection from using a `for` loop to fetch each URL and write the response body to a file one at a time for 2422 comics, 

```
$ time go run main.go

real    2m55.363s
user    0m2.743s
sys     0m2.174s
```
