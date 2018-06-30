// Parse a log file, and extracted the first column (space delemited)
// Then count the number of appearance.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read the string filename
	var filename string
	filename = "test7.txt"

	counts := make(map[string]int)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file!")
		os.Exit(2)
	}
	defer file.Close()

	input := bufio.NewScanner(file)

	// Scan per line
	for input.Scan() {
		// Split line by space and get the first one.
		word := strings.Fields(input.Text())[0]
		counts[word]++
	}
	for w, n := range counts {
		if n > 0 {
			fmt.Printf("%s\t%d\n", w, n)
		}
	}
}
