// your own advanced search version

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	// resp, err := http.Get(IssuesURL + "?q=" + q)
	// if err != nil {
	// 	return nil, err
	// }
	// //!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// ReadIssue read an issue
// GET /repos/:owner/:repo/issues/:number
// `:owner` mean you have to replace it with string variable
func ReadIssue(owner string, repo string, number string) (*Issue, error) {
	q := strings.Join([]string{
		APIURL, "repos", owner, repo, "issues", number,
	}, "/")

	req, err := http.NewRequest("GET", q, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return &result, nil
}

// CreateIssue create an issue
// POST /repos/:owner/:repo/issues
func CreateIssue(owner, repo, title string) (*Issue, error) {
	// url := strings.Join([]string{
	// 	APIURL, "repos", owner, repo, "issues",
	// }, "/")
	EDITOR := os.Getenv("EDITOR")
	if EDITOR == "" {
		fmt.Println("You haven't set the environment EDITOR")
		os.Exit(2)
	}
	tmpfile, err := ioutil.TempFile("", "githubIssue")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // Clean up
	editorCmd := exec.Command(EDITOR, tmpfile.Name())
	err = editorCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadFile(tmpfile.Name())

}

//!-
