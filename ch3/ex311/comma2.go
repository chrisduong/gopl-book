// “Enhance comma so that it deals correctly with floating-point numbers and an optional sign.”
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	b := bytes.Buffer{}
	// Need to find the mantissa (significant) of the number
	mantissaStart := 0
	// If the first char is sign, write it to buffer
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	// repeat the flow of comma.go with alteration but same effect
	mantissa := s[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3
	if pre > 0 {
		b.Write([]byte(mantissa[:pre]))
		if len(mantissa) > pre {
			b.WriteString(",")
		}
	}
	for i, c := range mantissa[pre:] {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	b.WriteString(s[mantissaEnd:])
	return b.String()
}
