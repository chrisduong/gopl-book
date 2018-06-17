// Modify issues to report the results in age categories, say less than a month old, less than a year old, and more than a year old
// Page: 365
//+
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chrisduong/gopl-book/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	fmt.Println("Issues less than a month")
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours()/24 <= 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	// Remember to exclude month-old issues for prettier
	fmt.Println("Issues less than a year")
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours()/24 > 30 &&
			time.Since(item.CreatedAt).Hours()/24 <= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("Issues more than a year")
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours()/24 > 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

//* go run main.go repo:golang/go is:open json decoder
