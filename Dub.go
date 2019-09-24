package main

import (

	"net/http"
	"github.com/creator/Dub/handlers"
)

// Go Gophers out there will need a starting point premade for them having solved a thousand problems.

func main() {

	//type localhost:1337/index in your browsers address bar
	http.HandleFunc("/login", handlers.LoginHandler)
	//parses the entire static folder and sub folders
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":1337", nil)
	//end of main function calling ListenAndServe to produce output to browser.
	//listening on port:1337 please open browser with link localhost:1337/index to view site.
}

