package handlers


import(
	"fmt"
	"html/template"
	"net/http"
	"github.com/Creator/Dub/static/goget"

)



//IndexHandler is the Handler for the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {

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

