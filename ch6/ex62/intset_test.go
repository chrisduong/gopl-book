package intset

import "testing"

func TestAddAllInset(t *testing.T) {
	var x IntSet
	x.AddAll(1, 144, 9)

	if x.String() != "{1 9 144}" {
		t.Log(x)
		t.Fail()
	}
}
