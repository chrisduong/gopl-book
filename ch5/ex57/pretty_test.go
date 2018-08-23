package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var input = `
<html>
<body>
	<p class="something" id="short"><span class="special">hi</span></p>
	<br>
	<!-- this is the comment -->
</body>
</html>
`

// But html.Parse parses pretty much anything, so this test is useless.
func TestPrettyOutputCanBeParsed(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}

	// You need to initialize a pointer because the method `Write` is on the pointer of `Buffer` type
	b := &bytes.Buffer{}
	var pp PrettyPrinter
	err = pp.Pretty(b, doc)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	// Check if the output of the method can be parsed
	_, err = html.Parse(bytes.NewReader(b.Bytes()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

// TestPrettyShortForm check if the Shortform can be done
func TestPrettyShortForm(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(input))

	b := &bytes.Buffer{}
	var pp PrettyPrinter
	err = pp.Pretty(b, doc)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "<br/>") {
			return
		}
	}

	t.Fail()
}
