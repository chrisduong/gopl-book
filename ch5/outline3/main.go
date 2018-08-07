// Modify outline, to check the traverse which ommit the Script/Style Node.
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
		fmt.Fprintf(os.Stderr, "outline special: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

	// NOTE: Recursion without loop
	// if n.FirstChild != nil {
	// 	outline(stack, n.FirstChild)
	// }

	// if n.NextSibling != nil {
	// 	outline(stack, n.NextSibling)
	// }
	// XXX: output look totally, because of different traverse way, which make the stack is bigger. it keeps accumlating the Child/Sibling nodes unitl exhaustion. Then it return to its caller.
	// [html]
	// [html head]
	// [html head meta]
	// [html head meta meta]
	// [html head meta meta meta]
	// [html head meta meta meta title]
	// [html head meta meta meta title link]
	// [html head body]

}

//!-
