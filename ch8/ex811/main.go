// “Following the approach of mirroredQuery in Section 8.4.4, implement a variant of fetch that requests several URLs concurrently. As soon as the first response arrives, cancel the other requests”

package main

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) { /* ... */ }
