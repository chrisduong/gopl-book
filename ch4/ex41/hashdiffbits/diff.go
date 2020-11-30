// Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section 2.6.2.) Page 134

package hashdiffbits

import (
	"crypto/sha256"

	"github.com/chrisduong/gopl-book/ch2/popcount"
)

// Diff compare 2 slice of byte
func Diff(a, b []byte) int {
	aSha := sha256.Sum256(a)
	bSha := sha256.Sum256(b)

	// XOR to find the different bit for each byte
	var xoredBoth [32]byte
	for i := range xoredBoth {
		xoredBoth[i] = aSha[i] ^ bSha[i]
	}
	// MORE: Can consider the solution to avoid loop through 32 bytes. SEE: https://stackoverflow.com/a/13657862 for "Convert [8]byte to a uint64"
	// Apply PopCount function for each byte
	var result int
	for i := range xoredBoth {
		result += popcount.PopCount(uint64(xoredBoth[i]))
	}
	return result
}
