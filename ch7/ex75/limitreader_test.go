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
		instance int
		bytes    []byte
		limit    int
		want     Want
	}{
		{1, make([]byte, 8), 9, Want{io.EOF, 8}},
		{2, make([]byte, 7), 9, Want{io.EOF, 7}},
		{3, make([]byte, 8), 4, Want{io.EOF, 4}},
	}

	for _, test := range tests {
		r := LimitReader(strings.NewReader(s), test.limit)
		n, err := r.Read(test.bytes)
		if (err != test.want.err) || (n != test.want.n) {
			t.Errorf("Instance %d: got %d, want %d", test.instance, n, test.want.n)
		}
	}

}
