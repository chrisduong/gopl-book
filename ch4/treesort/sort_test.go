// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"sort"
	"testing"

	"github.com/chrisduong/gopl.io/ch4/treesort"
)

func TestSort(t *testing.T) {
	// TEST: with non random Slice
	data := []int{7, 39, 40, 6, 5, 71}

	// data := make([]int, 50)
	// for i := range data {
	// 	data[i] = rand.Int() % 50
	// }
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
