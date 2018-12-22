// (*IntSet).UnionWith computes the union of two sets using |, the word-parallel bitwise OR operator. Implement methods for IntersectWith, DifferenceWith, and SymmetricDifference for the corresponding set operations. (The symmetric difference of two sets contains the elements present in one set or the other but not both.)”

package inset3

import "testing"

func TestIntersectWith(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	var y IntSet
	y.Add(1)
	y.Add(122)
	y.Add(9)

	x.IntersectWith(&y)

	if x.String() != "{1 9}" {
		t.Log(x.String())
		t.Fail()
	}

}
