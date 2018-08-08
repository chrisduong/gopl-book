// https://gist.github.com/Xeoncross/8bbb84bc4bf540bd907f79ee17c4e1fc
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
    body content
    <p>more <a href="">content</a></p>
    <p>This <a href="/foo"><em>important</em> link <br> to
    foo</a> is here</p>
    <p>Call at <a href="mailto:sam@example.com">sam@example.com</a></p>
    <div>
    <span>sam2@example.com</span></div>
    <p>Hello and <a href="">example.com</a>.</span></p>
    <em>are all valid. You can email "john" if you need me.</em>
</body>
</html>`

func main() {

	doc, err := html.Parse(strings.NewReader(htm))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find all the text nodes that are not children of a <p>
	matcher := func(node *html.Node) (keep bool, exit bool) {
		if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
			keep = true
		}
		if node.DataAtom == atom.P {
			exit = true
		}
		return
	}
	nodes := TraverseNode(doc, matcher)

	for i, node := range nodes {
		// fmt.Printf("Node Text: %s, Node Atom: %v \n", node.Data, node.DataAtom)

		fmt.Println(i, renderNode(node))
	}

}

// TraverseNode collecting the nodes that match the given function
func TraverseNode(doc *html.Node, matcher func(node *html.Node) (bool, bool)) (nodes []*html.Node) {
	var keep, exit bool
	var f func(*html.Node)
	f = func(n *html.Node) {
		keep, exit = matcher(n)
		if keep {
			nodes = append(nodes, n)
		}
		if exit {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}

// Works better than: https://github.com/yhat/scrape/blob/master/scrape.go#L129
// because you can cut the search short from the matcher function

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
