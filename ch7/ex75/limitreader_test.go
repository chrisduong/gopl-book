package limitreader

import (
	"io"
	"strings"
	"testing"
)

type Want struct {
	err error
	n   int
}

// TODO: fix auto update go.mod

func TestLimitReader(t *testing.T) {
	s := "hi there"
	var tests = []struct {
		bytes []byte
		limit int
		want  Want
	}{
		{make([]byte, 8), 9, Want{io.EOF, 8}},
		{make([]byte, 7), 9, Want{io.EOF, 7}},
		{make([]byte, 8), 4, Want{io.EOF, 5}},
	}

	// TODO: need to fix missing full Errorf message
	for _, test := range tests {
		r := LimitReader(strings.NewReader(s), test.limit)
		n, err := r.Read(test.bytes)
		if (err != test.want.err) && (n != test.want.n) {
			t.Errorf("%s: got %d, want %d", test.bytes, n, test.want.n)
			// t.Log(n)
			// t.Log(err)
			// t.Fail()
		}
	}

}
