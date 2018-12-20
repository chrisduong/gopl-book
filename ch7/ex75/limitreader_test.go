package limitreader

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hi there"
	b := make([]byte, 1024)
	r := LimitReader(strings.NewReader(s), 4)
	n, err := r.Read(b)
	if (err != nil) && (n != 4) {
		t.Log(n)
		t.Log(err)
		t.Fail()
	}
	// need to read again to return EOF
	n, err = r.Read(b)
	if err != io.EOF {
		t.Log(err)
		t.Fail()
	}

}
