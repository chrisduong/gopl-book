//+ Write a version of rotate that operates in a single pass

package rotate

// This is the final version which rotates in a single pass
func rotate(ints []int, n int) {
	// In case the number of rotates are larger than the length,
	// we need to reduce to the similar case to avoid out of bound.
	if n > len(ints) {
		rotate(ints, n-len(ints))
		return
	}
	for i, c := len(ints)-n, 0; c < len(ints); i, c = i-n, c+1 {
		// This is to deal with negative index which golang doesn't support
		if i < 0 {
			i = i + len(ints)
		}
		ints[0], ints[i] = ints[i], ints[0]
	}
}

// This function use copy() to rotate
func rotateV2(ints []int, n int) {
	if n > len(ints) {
		rotate(ints, n-len(ints))
		return
	}

	left := make([]int, n, n)
	copy(left, ints[:n])

	// Copy all the Right members
	copy(ints, ints[n:])

	// Copy all to the right from the rotation point
	copy(ints[len(ints)-n:], left)
}

func rotateV1(ints []int, n int) []int {
	var z []int
	var zlen int
	// Check if number of rotate is greater than the length
	if n > len(ints) {
		return rotateV1(ints, n-len(ints))
	}
	zlen = len(ints) + n
	if zlen <= cap(ints) {
		// There is room to expand the slice.
		z = ints[:zlen]
	} else {
		// There is insufficient space.
		// Grow by exactly needed
		zcap := zlen
		z = make([]int, zlen, zcap)
		copy(z, ints)
	}

	// We start j from the end of the new slice
	for i, j := n-1, zlen-1; i >= 0; i, j = i-1, j-1 {
		z[j] = z[i]
	}

	// We only care from rotate point n
	return z[n:]
}
