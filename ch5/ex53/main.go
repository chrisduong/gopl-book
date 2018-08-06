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
			return texts
		}
	}
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	// XXX: if you ommit sth on the traverse way, you should use the For loop to traverse, to make sure you don't miss the sibling node as the early return.

	// XXX: this code will miss the sibling of Script/Style Node
	// Traverse nodes
	// if n.FirstChild != nil {
	// 	texts = getText(texts, n.FirstChild)
	// }

	// if n.NextSibling != nil {
	// 	texts = getText(texts, n.NextSibling)
	// }

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = getText(texts, c)
	}

	return texts
}
