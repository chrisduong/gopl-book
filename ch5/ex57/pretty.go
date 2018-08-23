// Develop startElement and endElement into a general HTML pretty-printer. Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use short forms like <img/> instead of <img></img> when an element has no children. Write a test to ensure that the output can be parsed successfully. (See Chapter 11.)

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var pp PrettyPrinter
	// Print pretty output
	pp.Pretty(os.Stdout, doc)
}

// PrettyPrinter type contain io.Writer so that you can write to
// an `err` if the Writing causing any errors.
type PrettyPrinter struct {
	w   io.Writer
	err error
}

// Pretty print pretty output of the HTML node
func (pp PrettyPrinter) Pretty(w io.Writer, n *html.Node) error {
	pp.w = w
	pp.err = nil
	pp.forEachNode(n, pp.startElement, pp.endElement)
	return pp.err
}

// forEachNode traverse node and apply the logic
func (pp PrettyPrinter) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	// make sure to return if there is any error in pp's State
	if pp.err != nil {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pp.forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
	if pp.err != nil {
		return
	}
}

func (pp PrettyPrinter) printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(pp.w, format, args...)
	pp.err = err
}

// Initialize the depth to indent
var depth int

// start will print first appearance of each Element it meet
func (pp PrettyPrinter) startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		pp.printNode(n)
	case html.TextNode:
		pp.printText(n)
	case html.CommentNode:
		pp.printComment(n)
	}
}

// endElement print end Element Node, and reduce the depth
func (pp PrettyPrinter) endElement(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}
	depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
}

// startElement will print with 2 spaces indented for each new child element
func (pp PrettyPrinter) printNode(n *html.Node) {
	// We always start with Element Node
	// If the Node has child, the end character is always ">"
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}
	// If it is Element Node, pretty print its attributes
	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	pp.printf("%*s<%s%s%s\n", depth*2, "", n.Data, attrStr, end)
	depth++
}

// printText when meet the Text Node, it always has no child,
// means no need to increase the depth
func (pp PrettyPrinter) printText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	pp.printf("%*s%s\n", depth*2, "", n.Data)
}

func (pp PrettyPrinter) printComment(n *html.Node) {
	pp.printf("<!--%s-->\n", depth*2, "", n.Data)
}
