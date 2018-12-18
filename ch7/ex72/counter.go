// Write a function CountingWriter with the signature below that, given an io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer

package counter

import (
	"io"
)

type byteCounter struct {
	w     io.Writer
	count int64
}

func (c *byteCounter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.count += int64(n)
	return
}

// CountingWriter return an byteCounter struct and its count
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteCounter{w, 0}
	return c, &c.count
}

