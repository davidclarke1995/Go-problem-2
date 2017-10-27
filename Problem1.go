// Problem 7
// David Clarke
// Websites used https://stackoverflow.com/questions/3676127/how-do-i-make-a-text-input-non-editable
//https://stackoverflow.com/questions/19233415/how-to-make-type-number-to-positive-numbers-only
//https://stackoverflow.com/questions/12218670/how-to-increment-a-bootstrap-progress-bar

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
	Guess int
	check bool
	NewMessage string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Changes the font to Bold
	w.Header().Set("Content-Type","text/html")

	//Takes the HTML file "guess.html" and outputs it in the web app
	http.ServeFile(w, r, "guess.html")

}

func guessHandler(w http.ResponseWriter, r *http.Request){

		 message :="Guess a number between 1 and 20"

		 guess, _ := strconv.Atoi(r.FormValue("guess"))
		
			rand.Seed(time.Now().UTC().UnixNano())
		
			target:= rand.Intn(20-1)

			var cookie, err = r.Cookie("target")
		
			if err == nil{
				
				cookie = &http.Cookie{
				Name: "target",
				Value: strconv.Itoa(target),
				Expires: time.Now().Add(72 * time.Hour),
			}
			
			http.SetCookie(w,cookie)	
				
			}

			str := &myMsg{Message:message, Guess:guess, check: false}
		
			if guess == target{
				str.NewMessage = "Well done !! You guessed Correctly!"
				str.check = true
			}else if guess > target{
				str.NewMessage = "Try again, guess is too high"
			}else if guess < target{
				str.NewMessage = "Try again, guess is too low"
			}

			t, _ := template.ParseFiles("guess.tmpl")

			t.Execute(w, str)
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}