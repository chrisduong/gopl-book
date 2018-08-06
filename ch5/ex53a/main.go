package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// https://github.com/nasciiboy/TGPL-Exercises/blob/master/05-03/nasciiboy/main.go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	pPrint(doc)
}

// visit appends to links each link found in n and returns the result.
func pPrint(n *html.Node) {
	// TODO: why this won't skip Child and Sibling of Script Node.
	// Like footer Node. which is the Sibling.
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}

	if n.Type == html.TextNode {
		fmt.Printf("%s", n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pPrint(c)
	}
}
