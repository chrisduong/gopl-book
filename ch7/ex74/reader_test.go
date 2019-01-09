package reader

import (
	"testing"

	"golang.org/x/net/html"
)

func TestNewReaderWithHTML(t *testing.T) {
	s := "<html><body><p>hi</p></body></html>"
	doc, err := html.Parse(NewReader(s))
	if err != nil {
		t.Log(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
		t.Log(err)
		t.Fail()
	}

}
