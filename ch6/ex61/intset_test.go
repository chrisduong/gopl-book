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

func TestRemoveInset(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	x.Remove(9)

	if x.Has(9) || x.Len() != 2 {
		t.Log(x.String())
		t.Fail()
	}
}

func TestCopyInset(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y := x.Copy()

	if !y.Has(9) {
		t.Log(x.String())
		t.Log(y.String())
		t.Logf("%T: &x=%p", &x, &x)
		t.Log(&y)
		t.Fail()
	}
}
