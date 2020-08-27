// Excercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// Create a map to store Filename with Array of lines
	countFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, countFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countFiles map[string][]string) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !arrayContains(countFiles[text], name) {
			countFiles[text] = append(countFiles[text], name)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func arrayContains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

//!+ TEST
// go run ch1/ex1_4/main.go ch1/ex1_4/data1.txt ch1/ex1_4/data2.txt
//!+ OUTPUT
// 3       aaa     [ch1/ex1_4/data1.txt ch1/ex1_4/data2.txt]
// 2       abc     [ch1/ex1_4/data1.txt ch1/ex1_4/data2.txt]
