// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 77
// “Modify the Lissajous server to read parameter values from the URL.
// For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20
// sets the number of cycles to 20 instead of the default 5.
// Use the strconv.Atoi function to convert the string parameter into an integer.
// You can see its documentation with go doc strconv.Atoi.”

//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	var cycles = 5
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, cycles)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cyclesStr := r.FormValue("cycles")

	if cyclesStr != "" {
		cycles, err := strconv.Atoi(cyclesStr)
		if err != nil {
			fmt.Fprintf(w, "bad parameter for cycles: %s", cyclesStr, err)
			return
		}
		lissajous(w, cycles)
	}

}

// func lissajouse
// `cycles`: number of complete x oscillator revolutions
func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
//!+RUN
// 1. Run to export GIF file
// $ go run main > out.gif
// 2. Run as a web service
// $ go run main web
// open http://localhost:8000

//!-RUN
