// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Exercise 3.10
// See page 240.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+ No recurisve solution
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	r := n % 3
	if n/3 == 0 {
		buf.WriteString(s)
		return buf.String()
	}
	buf.WriteString(s[:r])
	for i := 0; i < n/3; i++ {
		buf.WriteByte(',')
		buf.WriteString(s[r : r+3])
		r += 3
	}
	return buf.String()
}

//!-
