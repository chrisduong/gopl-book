// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

// This array represents  the number of set bits from 0 to 255 (maximum number of 1 byte - 2**8-1)
// The next number's set bits will be " its 1 shifted Right number" (half index/value)
// + 1 if it is odd
func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] + // Count on the last byte (8bits),
		// byte to make sure it only capture the last byte after shifting.
		pc[byte(x>>(1*8))] + // Count on the second byte
		pc[byte(x>>(2*8))] + // So on to be 64bits = 8*8
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-
