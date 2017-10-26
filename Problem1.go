// Problem 4
// David Clarke
// Websites used https://stackoverflow.com/questions/2906582/how-to-create-an-html-button-that-acts-like-a-link

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
				
				target, _ = strconv.Atoi(cookie.Value)
				if target ==0{
					target = rand.Intn(20-1)
				}
			}
		
			cookie = &http.Cookie{
				Name: "target",
				Value: strconv.Itoa(target),
				Expires: time.Now().Add(72 * time.Hour),
			}
			
			http.SetCookie(w,cookie)
			
			t, _ := template.ParseFiles("guess.tmpl")

			t.Execute(w, &myMsg{Message:message})
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/guess", guessHandler)
}