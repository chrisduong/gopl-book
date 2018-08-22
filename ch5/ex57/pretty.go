// Develop startElement and endElement into a general HTML pretty-printer. Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use short forms like <img/> instead of <img></img> when an element has no children. Write a test to ensure that the output can be parsed successfully. (See Chapter 11.)

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// Print pretty output
	Pretty(os.Stdout, doc)
}

// Pretty print pretty output of the HTML node
func Pretty(w io.Writer, n *html.Node) error {
	// Comment nodes, text nodes and element nodes

	if n.Type == html.CommentNode {
		// if the Element Node does not have child
		// print short form
		if n.FirstChild == nil {
			fmt.Printf("<%s/>", n.Data)
		}
		// Print StartElement
		fmt.Printf("<%s", n.Data)

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Pretty(w, c)
	}
	return nil
}
