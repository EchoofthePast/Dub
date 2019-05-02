package handlers



import (
	"fmt"
	"net/http"
	"html"
	"strings"
	"html/template"
	"os"
)


func InputHandler(w http.ResponseWriter, r *http.Request) {

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
