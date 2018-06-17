// Modify charcount to count letters, digits, and so on in their Unicode categories, using functions like unicode.IsLetter.

// Page: 324
//+
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Unicode categories
const (
	CharIsSpace = iota
	CharIsSymbol
	CharIsMark
	CharIsDigit
	CharIsPrint
	CharIsPunct
	CharIsLetter
	CharIsNumber
	CharIsControl
	CharIsGraphic
)

//!+
func main() {
	var counts [10]int              // counts of Unicode categories which has max 10 length
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	var catname string              // Categorie name
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		// ReadRune() would convert invalid unicode char to ReplacementChar
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// Count every category
		switch {
		case unicode.IsSpace(r):
			counts[CharIsSpace]++
		case unicode.IsSymbol(r):
			counts[CharIsSymbol]++
		}
		utflen[n]++
	}
	fmt.Printf("rune categories\tcount\n")
	for i, n := range counts {
		if n != 0 {
			switch i {
			case CharIsSpace:
				catname = "Space"
			case CharIsSymbol:
				catname = "Symbol"
			}
			fmt.Printf("%s\t\t%d\n", catname, n)
		}
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
