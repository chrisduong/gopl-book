package intset

import "testing"

func TestLenInset(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	if x.Len() != 3 {
		t.Log(x.String())
		t.Fail()
	}
}
