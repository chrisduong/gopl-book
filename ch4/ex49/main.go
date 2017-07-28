// Write a program wordfreq to report the frequency of each word in an input text file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// NOTE: Don't know why no parameters is accepted for `bufio.ScanWords`.
	// MAYBE: input is already a Scanner so it is OK.
	// Split words in Scanner
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	for w, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, w)
		}
	}
}
