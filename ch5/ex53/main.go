// “Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into <script> or <style> elements, since their contents are not visible in a web browser.”

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
		os.Exit(1)

	}
	for _, text := range getText(nil, doc) {
		fmt.Printf("Text is %s \n", text)
	}
}

// printText print the contents of all text nodes in an HTML doc tree
func getText(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			// Do nothing
		}
	}
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	// Traverse nodes
	if n.FirstChild != nil {
		texts = getText(texts, n.FirstChild)
	}

	if n.NextSibling != nil {
		texts = getText(texts, n.NextSibling)
	}

	return texts
}
