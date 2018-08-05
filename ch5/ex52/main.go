// “Write a function to populate a mapping from element names—p, div, span, and so on—to the number of elements with that name in an HTML document tree”

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
	}

	elements := make(map[string]int)
	for k, v := range count(elements, doc) {
		fmt.Printf("Element %s: %d \n", k, v)
	}
}

// count() count elements found in n and returns the result.
func count(elements map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode && n.Data != "" {
		elements[n.Data]++
	}

	// Traverse nodes
	if n.FirstChild != nil {
		count(elements, n.FirstChild)
	}

	if n.NextSibling != nil {
		count(elements, n.NextSibling)
	}

	return elements
}
