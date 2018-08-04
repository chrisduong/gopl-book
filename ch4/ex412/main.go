// “The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.”

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Comic describe Comic JSON object
type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string // url
}

func getComic(n int) (Comic, error) {
	var commic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return commic, err
	}

	if resp.StatusCode != http.StatusOK {
		return commic, fmt.Errorf("Cannot get commic %d: %s", n, resp.Status)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&commic); err != nil {
		return commic, err
	}

	return commic, nil

}

func main() {

}
