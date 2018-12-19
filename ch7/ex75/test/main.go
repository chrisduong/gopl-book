// The LimitReader function in the io package accepts an io.Reader r and a number of bytes n, and returns another Reader that reads from r but reports an end-of-file condition after n bytes. Implement it
// “func LimitReader(r io.Reader, n int64) io.Reader”

package main

import (
	"fmt"
	"io"
	"strings"
)

type IOLimitReader struct {
	R io.Reader
	N int64
}

// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
func (l *IOLimitReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if l.N < int64(len(p)) {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	fmt.Println(p)
	return
}

// LimitReader returns a Reader that reads from r
// but stops with EOF after n bytes.
// The underlying implementation is a *LimitedReader.
func LimitReader(r io.Reader, n int64) *IOLimitReader {
	return &IOLimitReader{r, n}
}

func main() {
	s := "hi there"
	b := make([]byte, 1024)
	r := LimitReader(strings.NewReader(s), 4)
	_, _ = r.Read(b)
	n, err := r.Read(b)
	fmt.Printf("Bytes: %d, Err: %v", n, err)
}
