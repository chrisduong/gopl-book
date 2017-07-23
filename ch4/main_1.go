// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package main

import (
	"fmt"

	"github.com/chrisduong/gopl.io/ch4/treesort"
)

func main() {
	data := []int{7, 39, 40, 6, 5, 71}
	treesort.Sort(data)
	fmt.Printf("sorted: %v", data)
}

//!-
