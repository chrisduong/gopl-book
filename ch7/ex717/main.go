// “Extend xmlselect so that elements may be selected not just by name, but by their attributes too, in the manner of CSS, so that, for instance, an element like <div id="page" class="wide"> could be selected by a matching id or class as well as its name”

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("sample.xml", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	dec := xml.NewDecoder(f)
	var stack [][]string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			nameAndAttributes := make([]string, 0)
			nameAndAttributes = append(nameAndAttributes, tok.Name.Local)
			for _, attr := range tok.Attr {
				if attr.Name.Local == "id" || attr.Name.Local == "class" {
					nameAndAttributes = append(nameAndAttributes, attr.Value)
				}
			}
			stack = append(stack, nameAndAttributes)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				for _, nameAndAttributes := range stack {
					fmt.Printf("%s ", strings.Join(nameAndAttributes, "|"))
				}
				fmt.Printf(": %s\n", tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		for _, element := range x[0] {
			if element == y[0] {
				y = y[1:]
				break
			}
		}
		x = x[1:]
	}
	return false
}

//!-
// "args": ["div", "div", "h2"]
