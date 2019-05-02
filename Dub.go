package main

import (

	"html/template"
	"net/http"
	"github.com/echoofthepast/Dub/handlers"
)

// Go Gophers out there will need a starting point premade for them having solved a thousand problems.

func main() {

	//type localhost:1337/index in your browsers address bar
	http.HandleFunc("/index", handlers.IndexHandler)
	//type localhost:1337/input in your browsers address bar
	http.HandleFunc("/input", handlers.InputHandler)

	//parses the entire static folder and sub folders
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":1337", nil)
	//end of main function calling ListenAndServe to produce output to browser.
	//listening on port:1337 please open browser with link localhost:1337/index to view site.
}

//HandleHome is the Handler for the Home file
func HandleHome(w http.ResponseWriter, r *http.Request) {

	home := template.Must(template.ParseFiles("templates/home.html"))
	home.Execute(w, nil)

}
