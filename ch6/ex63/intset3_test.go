// (*IntSet).UnionWith computes the union of two sets using |, the word-parallel bitwise OR operator. Implement methods for IntersectWith, DifferenceWith, and SymmetricDifference for the corresponding set operations. (The symmetric difference of two sets contains the elements present in one set or the other but not both.)”

package inset3

import "testing"

func TestDifferentWith(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(14)
	x.Add(9)

	var y IntSet
	y.Add(1)
	y.Add(100)
	y.Add(9)
	y.Add(15)

	x.DifferentWith(&y)

	if x.String() != "{14}" {
		t.Log(x.String())
		t.Fail()
	}

}

func TestSymmetricDifference(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(14)
	x.Add(9)

	var y IntSet
	y.Add(1)
	y.Add(100)
	y.Add(9)
	y.Add(15)

	x.SymmetricDifference(&y)

	if x.String() != "{14 15 100}" {
		t.Log(x.String())
		t.Fail()
	}

}

func TestIntersectWith(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(14)
	x.Add(9)

	var y IntSet
	y.Add(1)
	y.Add(100)
	y.Add(9)
	y.Add(15)

	x.IntersectWith(&y)

	if x.String() != "{1 9}" {
		t.Log(x.String())
		t.Log(x)
		t.Fail()
	}

}
func TestUnionWith(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(9)

	var y IntSet
	y.Add(2)
	y.Add(122)
	y.Add(9)

	x.UnionWith(&y)

	if x.String() != "{1 2 9 122}" {
		t.Log(x.String())
		t.Fail()
	}

}
