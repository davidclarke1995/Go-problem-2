// Problem 5
// David Clarke
// Websites used https://stackoverflow.com/questions/3676127/how-do-i-make-a-text-input-non-editable
//https://stackoverflow.com/questions/19233415/how-to-make-type-number-to-positive-numbers-only
//https://stackoverflow.com/questions/16517718/bootstrap-number-validation

package main
//imports
import (
	"time"
	"net/http"
	"math/rand"
	"html/template"
	"strconv"
)

type myMsg struct {
    Message string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Changes the font to Bold
	w.Header().Set("Content-Type","text/html")

	//Takes the HTML file "guess.html" and outputs it in the web app
	http.ServeFile(w, r, "guess.html")

}

func guessHandler(w http.ResponseWriter, r *http.Request){

		 message :="Guess a number between 1 and 20"
		
			rand.Seed(time.Now().UTC().UnixNano())
		
			target:=0
			var cookie, err = r.Cookie("target")
		
			if err == nil{
				
				cookie = &http.Cookie{
				Name: "target",
				Value: strconv.Itoa(target),
				Expires: time.Now().Add(72 * time.Hour),
			}
			
			http.SetCookie(w,cookie)	
				
			}
		
			t, _ := template.ParseFiles("guess.tmpl")

			t.Execute(w, &myMsg{Message:message})
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)

	http.ListenAndServe(":8080", nil)
}