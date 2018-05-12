package main

import (
	"fmt"
	"strconv"
)

type stringer interface {
	String() string
}

// Sprint function to print out the type of variable interface
func Sprint(x interface{}) string {
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
		// ...similar cases for int16, uint32, and so on ...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map. pointer, slice, struct
		return "???"
	}
}

func main() {
	// TODO: need to Print the presentation of an Interface also
	var i = 16
	fmt.Printf("The String presentation of the integer i: %q", Sprint(i))
}
