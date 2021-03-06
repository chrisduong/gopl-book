// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 62. Ex 1.8.
// Modify fetch to add the prefix http:// to each argument URL
// if it is missing. You might want to use strings.HasPrefix

//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			resp.Body.Close()
			os.Exit(1)
		}
	}
}

//!-

//!+ TEST
// go run ch1/ex1_8/main.go http://gopl.io
// go run ch1/ex1_8/main.go gopl.io

//!-
