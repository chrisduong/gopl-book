// Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)

// See page 168.

package popcount2

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var result byte

	for i := uint(0); i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}

	return int(result)
}
