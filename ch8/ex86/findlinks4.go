// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string) // lists of URLs, may have duplicates
	// NOTE: try to use bufferred channel to limit to 3 links
	unseenLinks := make(chan string, 3) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

//!-

// ❯ go run ch8/ex86/findlinks4.go http://gopl.io/
// http://gopl.io/
// http://www.amazon.com/dp/020161586X?tracking_id=disfordig-20
// http://www.gopl.io/reviews.html
// http://golang.org/lib/godoc/analysis/help.html
// http://www.gopl.io/ch1.pdf
// http://www.barnesandnoble.com/w/1121601944
// https://github.com/golang/tools/blob/master/refactor/eg/eg.go
// http://www.informit.com/store/go-programming-language-9780134190440
// http://www.gopl.io/translations.html
// https://github.com/adonovan/gopl.io/
// http://www.amazon.com/dp/0134190440
// http://www.gopl.io/errata.html
// http://golang.org/s/oracle-user-manual
// https://github.com/golang/tools/blob/master/refactor/rename/rename.go
// http://www.amazon.com/dp/0131103628?tracking_id=disfordig-20
// http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/product-reviews/0134190440/ref=cm_cr_dp_see_all_summary
// http://www.infoq.com/articles/the-go-programming-language-book-review
// http://eli.thegreenplace.net/2016/book-review-the-go-programming-language-by-alan-donovan-and-brian-kernighan
// http://www.computingreviews.com/index_dynamic.cfm?CFID=15675338&CFTOKEN=37047869
// http://www.onebigfluke.com/2016/03/book-review-go-programming-language.html
// http://lpar.ath0.com/2015/12/03/review-go-programming-language-book
// https://www.usenix.org/system/files/login/articles/login_dec15_17_books.pdf
// http://www.acornpub.co.kr/book/go-programming
// http://www.williamspublishing.com/Books/978-5-8459-2051-5.html
// http://helion.pl/ksiazki/jezyk-go-poznaj-i-programuj-alan-a-a-donovan-brian-w-kernighan,jgopop.htm
// http://helion.pl/
// http://www.amazon.co.jp/exec/obidos/ASIN/4621300253
// http://www.maruzen.co.jp/corp/en/services/publishing.html
// https://novatec.com.br/livros/linguagem-de-programacao-go/
// http://novatec.com.br/
// https://www.tenlong.com.tw/products/9789864761333
// http://www.gotop.com.tw/
// https://www.amazon.cn/dp/B072LCX9S7
// http://www.pearsonapac.com/
// https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md
// https://golang.org/
// https://golang.org/lib/godoc/analysis/help.html
// https://golang.org/doc/
// https://golang.org/pkg/
// https://golang.org/project/
// https://golang.org/help/
// https://golang.org/blog/
// http://play.golang.org/
// http://golang.org/ref/spec#Method_sets
// https://go.googlesource.com/tools/+/master/godoc/analysis/README
// https://developers.google.com/site-policies#restrictions
// https://golang.org/LICENSE
