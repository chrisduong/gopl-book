// “Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("URL %s has %d words and %d images", url, words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	fmt.Println(words)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// TODO: the words supposed to be accummulated during recursion, but it won't
	if n.Type == html.TextNode {
		if strings.TrimSpace(n.Data) != "" {
			words += wordCount(n.Data)
		}
	}

	if n.Type == html.ElementNode {
		if n.Data == "img" {
			for _, i := range n.Attr {
				if i.Key == "src" {
					images++
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c)
	}

	return
}

func wordCount(s string) int {
	var count int
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)

	for input.Scan() {
		count++
	}

	return count
}
