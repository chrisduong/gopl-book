package main

import (
	"fmt"
	"os"
)

// NOTE: pitfall
func main() {
	tempDirs := []string{"1", "2", "3"}
	var rmdirs []func()
	for _, dir := range tempDirs {
		_ = os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir) // NOTE: incorrect!
			// This will print 3, 3, 3 on runtime
			fmt.Println(dir)

		})
	}
	for _, rmdir := range rmdirs {
		rmdir() // clean up
		// Only remove folder 3
	}

}
