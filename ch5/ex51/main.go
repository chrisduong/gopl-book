// Change the findlinks program to traverse the n.FirstChild linked list using recursive calls to visit instead of a loop

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// Check if FirstChild has any links then go to its' sibling
	// NOTE: https://www.w3.org/TR/dom/#nodes
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

// ! ./fetch https://golang.org | ./ex51
// /
// /
// #
// /doc/
// /pkg/
// /project/
// /help/
// /blog/
// http://play.golang.org/
// #
// #
// //tour.golang.org/
// /dl/
// //blog.golang.org/
// https://developers.google.com/site-policies#restrictions
// /LICENSE
// /doc/tos.html
