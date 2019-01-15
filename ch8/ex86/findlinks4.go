// ex8.6 is a depth-limited web crawler.
//
// Use a WaitGroup to determine when the work is done, the `tokens` chan as a
// semaphore to limit concurrent requests, and a mutex around the `seen` map to
// avoid concurrent reads and writes.
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"gopl.io/ch5/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		// Go to the next depth
		go crawl(link, depth+1, wg)
	}
}

func main() {
	flag.IntVar(&maxDepth, "d", 3, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0, wg)
	}
	wg.Wait()
}

// ❯ go run ch8/ex86a/main.go http://gopl.io/
// 0 http://gopl.io/
// 1 http://www.amazon.com/dp/020161586X?tracking_id=disfordig-20
// 1 http://www.barnesandnoble.com/w/1121601944
// 1 http://www.informit.com/store/go-programming-language-9780134190440
// 1 http://www.gopl.io/errata.html
// 1 http://www.gopl.io/ch1.pdf
// 1 http://www.amazon.com/dp/0131103628?tracking_id=disfordig-20
// 1 https://github.com/golang/tools/blob/master/refactor/eg/eg.go
// 1 http://www.gopl.io/reviews.html
// 1 http://www.gopl.io/translations.html
// 1 http://golang.org/s/oracle-user-manual
// 1 http://www.amazon.com/dp/0134190440
// 1 https://github.com/golang/tools/blob/master/refactor/rename/rename.go
// 1 http://golang.org/lib/godoc/analysis/help.html
// 1 https://github.com/adonovan/gopl.io/
// 2 https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md
// 2 http://www.acornpub.co.kr/book/go-programming
// 2 http://www.williamspublishing.com/Books/978-5-8459-2051-5.html
// 2 http://novatec.com.br/
// 2 http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/product-reviews/0134190440/ref=cm_cr_dp_see_all_summary
// 2 http://www.amazon.co.jp/exec/obidos/ASIN/4621300253
// 2 https://novatec.com.br/livros/linguagem-de-programacao-go/
// 2 http://www.infoq.com/articles/the-go-programming-language-book-review
// 2 http://www.computingreviews.com/index_dynamic.cfm?CFID=15675338&CFTOKEN=37047869
// 2 http://helion.pl/
// 2 http://www.onebigfluke.com/2016/03/book-review-go-programming-language.html
// 2 http://eli.thegreenplace.net/2016/book-review-the-go-programming-language-by-alan-donovan-and-brian-kernighan
// 2 http://www.maruzen.co.jp/corp/en/services/publishing.html
// 2 http://www.pearsonapac.com/
// 2 https://www.usenix.org/system/files/login/articles/login_dec15_17_books.pdf
// 2 https://www.tenlong.com.tw/products/9789864761333
// 2 http://lpar.ath0.com/2015/12/03/review-go-programming-language-book
// 2 https://www.amazon.cn/dp/B072LCX9S7
// 2 http://www.gotop.com.tw/
// 2 http://helion.pl/ksiazki/jezyk-go-poznaj-i-programuj-alan-a-a-donovan-brian-w-kernighan,jgopop.htm
// 2 http://www.google.com/intl/en/policies/privacy/
// 2 https://golang.org/blog/
// 2 https://golang.org/doc/
// 2 https://golang.org/help/
// 2 http://play.golang.org/
// 2 https://golang.org/doc/tos.html
// 2 https://golang.org/lib/godoc/analysis/help.html
// 2 https://go.googlesource.com/tools/+/master/godoc/analysis/README
// 2 https://golang.org/
// 2 https://golang.org/pkg/
// 2 https://golang.org/project/
// 2 http://golang.org/ref/spec#Method_sets
// 2 https://developers.google.com/site-policies#restrictions
// 2 https://golang.org/LICENSE
// 3 http://www.acornpub.co.kr/book/go-programming#header
// 3 http://www.acornpub.co.kr/about/map
// 3 http://www.yes24.com/24/goods/24334905
// 3 http://www.acornpub.co.kr/book/all
// 3 http://www.acornpub.co.kr/about/profile
// 3 http://www.aladin.co.kr/shop/wproduct.aspx?ItemId=76703559
// 3 http://www.acornpub.co.kr/contact/recruit
// 3 http://www.acornpub.co.kr/contact/errata
// 3 http://www.acornpub.co.kr/book/steadyseller
// 3 http://www.acornpub.co.kr/book/spring-webdev-set
