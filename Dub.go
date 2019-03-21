package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/echoofthepast/Dub/static/goget"
)


// Go Gophers out there will need a starting point premade for them having solved a thousand problems.

func main() {

	//type localhost:1337/home in your browsers address bar
	http.HandleFunc("/home", HandleHome)

	//type localhost:1337/index in your browsers address bar
	http.HandleFunc("/index", HandleIndex)

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

	data = &goget.DataStruct{

		Name:    name,
		Surname: surname,
		Both:    name + surname,
		Url:     "",
		Image1:  image1,
		Image2:  image2,
		Info:    info,

	}

	return data

}

func inputHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var namestring string
	var snamestring string
	var infostring string
	var medicationstring string

	fname := r.Form["fname"] //
	sname := r.Form["sname"]
	info := r.Form["info"]
	medication := r.Form["medication"]

	for index := 0; index > 1; index++ {

		namestring = fname[index]
		snamestring = sname[index]
		infostring = info[index]
		medicationstring = medication[index]

	}

	name := namestring
	surname := snamestring
	information := infostring
	meds := medicationstring

	fmt.Println("Name", name)
	fmt.Println("Surname", surname)
	fmt.Println("Information", information)
	fmt.Println("Medication", meds)

	//simply writing to file in the same format.

	tmpl := template.Must(template.ParseFiles("templates/input.html"))
	tmpl.Execute(w, nil)

}

//HandleIndex is the Handler for the index file
func HandleIndex(w http.ResponseWriter, r *http.Request) {

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
