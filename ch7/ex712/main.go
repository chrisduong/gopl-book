// “Change the handler for /list to print its output as an HTML table, not text. You may find the html/template package (§4.6) useful.”

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var dbMux sync.Mutex

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong price value: %q\n", priceStr)
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	dbMux.Lock()
	db[item] = dollars(price)
	dbMux.Unlock()
	fmt.Fprintf(w, "Update item %s's price to %s\n", item, db[item])
}

// templateData contains a list of items
type templateData struct {
	ItemMap database
}

var itemList = template.Must(template.New("itemlist").Parse(`
<h1> items </h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $item, $price := .ItemMap}}
<tr>
  <td>{{$item}}</td>
  <td>{{$price}}</td>
</tr>
{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := itemList.Execute(w, &templateData{db}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to execute template: %q\n", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		// return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/update", http.HandlerFunc(db.update))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
