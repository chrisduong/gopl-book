// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 62.
// “The function call io.Copy(dst, src) reads from src and writes to dst.
// Use it instead of ioutil.ReadAll to copy the response body to os.Stdout
// without requiring a buffer large enough to hold the entire stream.
// Be sure to check the error result of io.Copy”

//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()

	}
}

//!-

//!+ TEST
// go run ch1/ex1_7/main.go https://golang.org

//!-
