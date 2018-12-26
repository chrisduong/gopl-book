// “Use the html/template package (§4.6) to replace printTracks with a function that displays the tracks as an HTML table. Use the solution to the previous exercise to arrange that each click on a column head makes an HTTP request to sort the table.”

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var trackTable = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
	  <meta charset="utf-8">
		<style media="screen" type="text/css">
		  table {
				border-collapse: collapse;
				border-spacing: 0px;
			}
		  table, th, td {
				padding: 5px;
				border: 1px solid black;
			}
		</style>
	</head>
	<body>
		<h1>Tracks</h1>
		<table>
		  <thead>
				<tr>
					<th><a href="/?sort=Title">Title</a></th>
					<th><a href="/?sort=Artist">Artist</a></th>
					<th><a href="/?sort=Album">Album</a></th>
					<th><a href="/?sort=Year">Year</a></th>
					<th><a href="/?sort=Length">Length</a></th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!+printTracks
//!+printTracks
func printTracks(writer io.Writer, tracks []*Track) {
	if err := trackTable.Execute(writer, tracks); err != nil {
		log.Fatal(err)
	}
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	sortBy := request.FormValue("sort")

	sort.Sort(clickedSort{tracks, func(track1, track2 *Track) bool {
		switch sortBy {
		case "Title":
			return track1.Title < track2.Title
		case "Year":
			return track1.Year < track2.Year
		case "Length":
			return track1.Length < track2.Length
		case "Artist":
			return track1.Artist < track2.Artist
		case "Album":
			return track1.Album < track2.Album
		}
		return false
	}})
	printTracks(responseWriter, tracks)
}

// clickSort contain what columns had been clicked
type clickedSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// Defines methods to satisfy the sort Interface
func (x clickedSort) Len() int           { return len(x.t) }
func (x clickedSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x clickedSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

