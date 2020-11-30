// Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section 2.6.2.) Page 134

package hashdiffbits_test

import (
	"testing"

	"github.com/chrisduong/gopl-book/ch4/ex41/hashdiffbits"
)

func TestHashDiffBits(t *testing.T) {
	diff_count := hashdiffbits.Diff([]byte("x"), []byte("X"))

	if diff_count != 125 {
		t.Logf("%d != 4", diff_count)
		t.Fail()
	}
}
