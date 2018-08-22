// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-

// ## Outputs
// cat gopl.html| go run ch5/outline/main.go

// ```console
// [html]
// [html head]
// [html head meta]
// [html head title]
// [html head script]
// [html head link]
// [html head style]
// [html body]
// [html body table]
// [html body table tbody]
// [html body table tbody tr]
// [html body table tbody tr td]
// ## NOTEs

// - You don't record the end node. For e.g. `</head>`
