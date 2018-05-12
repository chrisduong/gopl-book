// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Package charcount computes counts of Unicode characters.
package charcount

// CharCount return the counts of Unicide character
func CharCount(s string) map[rune]int {
	counts := make(map[rune]int) // counts of String (Unicode) characters
	// var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	// invalid := 0                    // count of invalid UTF-8 characters

	// TODO: ignore invalid unicode character as it cannot be shown
	for _, c := range s {
		counts[c]++
	}
	return counts
}

//!-
