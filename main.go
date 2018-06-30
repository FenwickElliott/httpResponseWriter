package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	endpoint := "http://example.com"
	resp, err := http.Get(endpoint)
	check(err)

	f, err := os.Create("endpoint")
	check(err)
	defer f.Close()

	err = resp.Write(f)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
