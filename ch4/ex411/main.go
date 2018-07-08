// Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required
// Page: 364
//+

// Using http package to connect to Github

package main

import (
	"fmt"
	"os"

	github "github.com/chrisduong/gopl-book/ch4/github"
)

// // CreateIssue create a new issue
// func CreateIssue(title string, body string) {

// }

// // GetIssue get an issue by its number
// func GetIssue(number int) {

// }

var usage = `usage:
search QUERY
[read|edit|close|open] OWNER REPO ISSUE_NUMBER
`

func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usageDie()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		github.SearchIssues(args)
		os.Exit(0)
	}
}
