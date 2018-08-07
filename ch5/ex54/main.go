// “Extend the visit function so that it extracts other kinds of links from the document, such as images, scripts, and style sheets.”

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks4: %v\n", err)
		os.Exit(1)

	}
	for _, text := range findlinks4(nil, doc) {
		fmt.Printf("Text is %s \n", text)
	}
}

// findlinks4 extract images and scripts text from HTML document
func findlinks4(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "img" || n.Data == "script" {
			for _, link := range n.Attr {
				if link.Key == "src" {
					links = append(links, link.Val)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = findlinks4(links, c)
	}
	return links
}
