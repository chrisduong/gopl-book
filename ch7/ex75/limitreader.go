// The LimitReader function in the io package accepts an io.Reader r and a number of bytes n, and returns another Reader that reads from r but reports an end-of-file condition after n bytes. Implement it
// “func LimitReader(r io.Reader, n int) *IOLimitReader  ”

package limitreader

import (
	"io"
)

type IOLimitReader struct {
	r io.Reader
	n int
}

func (l *IOLimitReader) Read(p []byte) (n int, err error) {
	if len(p) >= l.n {
		// Read up to the limit
		n, err = l.r.Read(p[:l.n])
	} else {
		// Read up to len(p) if p is shorter than the limit
		n, err = l.r.Read(p)
	}

	// If can read to n limit bytes or read up to len(p)
	if n == l.n || n == len(p) {
		err = io.EOF
	}
	return
}

// LimitReader returns a Reader that reads from r
// but stops with EOF after n bytes.
// The underlying implementation is a *LimitedReader.
func LimitReader(r io.Reader, n int) *IOLimitReader {
	return &IOLimitReader{r, n}
}
