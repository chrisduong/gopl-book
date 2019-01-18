// “Following the approach of mirroredQuery in Section 8.4.4, implement a variant of fetch that requests several URLs concurrently. As soon as the first response arrives, cancel the other requests”

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	// Get all URLs from arguments
	flag.Parse()
	urls := flag.Args()

	var wg sync.WaitGroup
	var abort = make(chan struct{})
	var responses = make(chan *http.Response)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequest("HEAD", url, nil)
			if err != nil {
				log.Printf("HEAD %s: %s", url, err)
			}
			req.Cancel = abort
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Failed to get response: %s", err)
				return
			}
			responses <- resp
		}(url)
	}
	// Fetch the latest response
	resp := <-responses
	defer resp.Body.Close()
	// Close other requests
	close(abort)

	fmt.Println(resp.Request.URL)
	for name, vals := range resp.Header {
		fmt.Printf("%s: %s\n", name, strings.Join(vals, ","))
	}
	wg.Wait()

}

// ❯ go run ch8/ex811/main.go http://gopl.io https://google.com
// http://www.gopl.io/
// Vary: Accept-Encoding
// Last-Modified: Thu, 08 Jun 2017 01:27:58 GMT
// Etag: "5938a81e-103a"
// Expires: Fri, 18 Jan 2019 04:50:59 GMT
// Content-Length: 4154
// X-Served-By: cache-sin18029-SIN
// X-Cache-Hits: 0
// Content-Type: text/html; charset=utf-8
// Accept-Ranges: bytes
// Age: 0
// Connection: keep-alive
// Server: GitHub.com
// Access-Control-Allow-Origin: *
// X-Github-Request-Id: 37CE:5F86:AC1A6:D2662:5C4158DB
// X-Cache: MISS
// Cache-Control: max-age=600
// Date: Fri, 18 Jan 2019 04:40:59 GMT
// Via: 1.1 varnish
// X-Timer: S1547786460.704625,VS0,VE251
// X-Fastly-Request-Id: 1358c65b0c156eca4883b8ffda30d4a5107405aa
// 2019/01/18 12:57:14 Failed to get response: Head http://gopl.io: net/http: request canceled
