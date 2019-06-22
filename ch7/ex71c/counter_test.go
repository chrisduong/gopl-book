package counter

import "testing"

func TestLineCounter(t *testing.T) {
	var lc LineCounter
	p := []byte(" one\ntwo\n three \n")
	_, _ = lc.Write(p)

	if lc != 3 {
		t.Logf("lines: %d != 3", lc)
	}
}

func TestWordCounter(t *testing.T) {
	var c WordCounter
	data := [][]byte{
		[]byte("The upcoming word is sp"),
		[]byte("lit across the buffer boundary. "),
		[]byte(" And this one ends on the buffer boundary."),
		[]byte(" Last words."),
	}
	for _, p := range data {
		_, _ = c.Write(p)
		// 	if n != len(p) || err != nil {
		// 		t.Logf(`bad write: p="%s" n=%d err="%s"`, string(p), n, err)
		// 		t.Fail()
		// }
	}
	if c != 20 {
		t.Logf("words: %d != 20", c)
		t.Fail()
	}
}
