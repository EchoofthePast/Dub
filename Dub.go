package main

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"os"
	"strings"

	//"time"

	"github.com/echoofthepast/Dub/static/goget"
)

// Go Gophers out there will need a starting point premade for them having solved a thousand problems.

func main() {

	//type localhost:1337/index in your browsers address bar
	http.HandleFunc("/index", indexHandler)
	//type localhost:1337/input in your browsers address bar
	http.HandleFunc("/input", inputHandler)

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

func blank() (data *goget.DataStruct) {

	image1 := "static/images/Gopher.png"
	image2 := "static/images/The Universe/The Universe.jpg"
	name := "Gopher"
	surname := "Google"
	info := "The one who does the work is the Gopher"
	id := "2000010101"

	data = &goget.DataStruct{

		Name:    name,
		Surname: surname,
		Both:    name + surname,
		Image1:  image1,
		Image2:  image2,
		Info:    info,
		ID: 	 id,
	}

	return data

}

func inputHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if r.Method == "GET" {
		fmt.Fprintf(os.Stdout, "GET, %q", html.EscapeString(r.URL.Path))
	} else if r.Method == "POST" {
		fmt.Fprintf(os.Stdout, "POST, %q", html.EscapeString(r.URL.Path))
	} else {
		http.Error(w, "Invalid request method.", 405)
	}

	method := r.Method

	fmt.Println(" Method:", method)

	var key string

	if r.Method == "POST" {

		for k, v := range r.Form {
			key = k
			fmt.Println("key:", key)
			fmt.Println("val:", strings.Join(v, ""))
		}

	} //end of if statement

	tmpl := template.Must(template.ParseFiles("templates/input.html"))
	tmpl.Execute(w, nil)

}

var call int

//animation engine

//HandleIndex is the Handler for the index file
func indexHandler(w http.ResponseWriter, r *http.Request) {

	var data *goget.DataStruct

	r.ParseForm()

	data = blank()

	//input from FrontEnd
	search := r.Form["searchbox"] //

	///indexsearch(search)

	for index := range search {

		if search[index] != " " {

			input := search[index]

			data = goget.Gopher(input)
			//
			for index := range input {
				//
				var letters [255]string
				letters[index] = string(input[index])

				var i int
				if letters[index] == " " {

					fmt.Println("Search Word ", index, letters[index])
					i++
					length := index
					fmt.Println("The Length of the Word is ", length)

				}

			}

		}

	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)

}
