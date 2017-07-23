// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import "fmt"

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		fmt.Printf("LoopAdd to root a value: %v\n", v)
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	fmt.Printf("*1 Add value %v to tree might be nil\n", value)
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		fmt.Printf("*2 create new subtree is %v\n", value)
		return t
	}
	fmt.Printf("*3 Tree %v is not nil\n", t.value)

	if value < t.value {
		fmt.Printf("*4a Add left of tree %v a value %v ", t.value, value)
		fmt.Println(". You have to go back to the parent tree later")
		t.left = add(t.left, value)
	} else {
		fmt.Printf("*4b add right of tree %v a value %v", t.value, value)
		fmt.Println(". You have to go back to the parent tree later")
		t.right = add(t.right, value)
	}
	fmt.Printf("*5 Go back to the parent tree %v\n", t.value)
	return t
}

//!-
