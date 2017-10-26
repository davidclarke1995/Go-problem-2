// Problem 2
// David Clarke
// Websites used https://golang.org/doc/articles/wiki/

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Changes the font to Bold
	w.Header().Set("Content-Type","text/html")

	//Output "Guessing Game" to browser
	fmt.Fprintln(w, "<h1>Guessing Game</h1>")

}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}