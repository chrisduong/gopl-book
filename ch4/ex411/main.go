// Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required
// Page: 364
//+

// Using http package to connect to Github

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// search return issues based on search term
func search(query []string) {
	result, err := SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range result.Items {
		format := "#%-5d %9.9s %.55s\n"
		fmt.Printf(format, issue.Number, issue.User.Login, issue.Title)
	}
}

// ReadIssue read an issue
func ReadIssue(owner string, repo string, number string) (*Issue, error) {
	q := url.QueryEscape(strings.Join([]string{GithubAPIUrl, "repos"}, "/"))
	req, err := http.NewRequest("GET", APIURL, nil)
	return nil, nil
}

// CreateIssue create a new issue
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
		fmt.Println(len(os.Args))
		usageDie()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		search(args)
		os.Exit(0)
	}

	if cmd == "read" {
		owner := args[1]
		repo := args[2]
		issueNumber := args[3]
		ReadIssue(owner, repo, issueNumber)
	}
}
