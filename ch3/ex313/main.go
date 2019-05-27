// “Write const declarations for KB, MB, up through YB as compactly as you can”

// See page 252

// Printints demonstrates the use of bytes.Buffer to format a string.
package main

const (
	KB = 1000
	MB = KB * KB // 1_000_000
	GB = MB * KB // 1_000_000_000
)
