// Problem 1
// David Clarke

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Browser renders html tags
	w.Header().Set("Content-Type","text/html");

	//Output "Guessing Game" to browser
	fmt.Fprintln(w, "Guessing Game")
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}