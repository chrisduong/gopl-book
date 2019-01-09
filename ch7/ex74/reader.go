// The strings.NewReader function returns a value that satisfies the io.Reader interface (and others) by reading from its argument, a string. Implement a simple version of NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string

package reader

import "io"

type StringReader string

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = len(*r)
	copy(p, []byte(*r))
	err = io.EOF // Must set EOF, otherwise it does not end
	return
}

func NewReader(str string) *StringReader {
	var s StringReader
	s = StringReader(str)
	return &s
}
