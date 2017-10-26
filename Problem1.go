// Problem 2
// David Clarke
// Websites used https://getbootstrap.com/docs/4.0/getting-started/introduction/#starter-template

package main

import (
//	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Changes the font to Bold
	w.Header().Set("Content-Type","text/html")

	//Takes the HTML file "problem3.html" and outputs it in the web app
	http.ServeFile(w, r, "problem3.html")

}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}