package main

import (
	"fmt"
)

func f() *int {
	v := 1
	return &v
}

func main() {
	var p = f()

	fmt.Println(p)          // Print address of local variable v
	fmt.Println(f() == f()) // Print false, as everytime function f invoke,
	// it returns new local variable's address
}
