// Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section 2.6.2.) Page 134

package hashdiffbits_test

import (
	"crypto/sha256"
	"testing"

	"github.com/chrisduong/gopl-book/ch4/hashdiffbits"
)

func TestHashDiffBits(t *testing.T) {
	sha1 := sha256.Sum256([]byte("x"))
	sha2 := sha256.Sum256([]byte("Y"))
	result := hashdiffbits.Diff(sha1, sha2)
	if result != 4 {
		t.Logf("%d != 4", result)
		t.Fail()
	}
}
