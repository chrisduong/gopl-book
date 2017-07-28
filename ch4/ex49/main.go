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
	// NOTE: See this example from bufio package to know the use case of ScanWords:
	//* https://golang.org/pkg/bufio/#example_Scanner_words
	// XXX: Read EmptyFinalToken https://play.golang.org/p/IlLOzWfThf to understand the way override default split function `func(data []byte, atEOF bool) (advance int, token []byte, err error)`
	//-- Return the index of the rest tokens, the next token, and error if any.

	// Apply `ScanWords function` on Split
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
