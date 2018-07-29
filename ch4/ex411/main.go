// Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required
// Page: 364
//+

// Using http package to connect to Github

package main

import (
	"fmt"
	"log"
	"os"
)

// search return issues based on search term
func search(query []string) {
	result, err := SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range result.Items {
		fmt.Printf(FORMAT, issue.Number, issue.User.Login, issue.Title)
	}
}

// read return the issue based on its number
func read(owner string, repo string, number string) {
	result, err := ReadIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(FORMAT, result.Number, result.User.Login, result.Title)
}

// create a new issue with a title, open text editor to input their own content
// TODO: find the way to call an EDITOR write a temp file then read. SEE: https://stackoverflow.com/a/6309753/1177314, https://gobyexample.com/spawning-processes
func create(owner, repo, title string) {
	issue := Issue{}
	fmt.Printf(FORMAT, issue.Number, issue.User.Login, issue.Title)
}

var usage = `usage:
search QUERY
create OWNER REPO
[read|update|delete] OWNER REPO ISSUE_NUMBER
`

func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(len(os.Args)) /**/
		usageDie()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		search(args)
		os.Exit(0)
	}

	if cmd == "read" {
		owner := args[0]
		repo := args[1]
		issueNumber := args[2]
		read(owner, repo, issueNumber)
	}

	if cmd == "create" {
		owner := args[0]
		repo := args[1]
		create(owner, repo)
	}
}
