package handlers

import(
	"fmt"
	"net/http"
	"html/template"
	"github.com/Creator/Dub/static/goget"
)

//DarknetHandler  handler function
func DarknetHandler(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm()
	data = blank()
	
	r.ParseForm()
	data = blank()
	
	method := r.Method
	
	fmt.Println(" Method:", method)
	
	search := r.Form["searchbox"] //
	
	for index := range search {
		
		if search[index] != "" {
			
			input := search[index]
			
			//Addition user error prevention required
			//data = gopher(input) //
			data = goget.Gopher(input)
			
			for index := range input {
				//
				var letters [255]string
				memory := len(letters)
				letters[index] = string(input[index])
				
				var i int
				var length1 int
				var length2 int

				if letters[index] == " " {
					
					fmt.Println("Search Word 1 ", index, letters[index])
					i++
					length1 = index
					//fmt.Println("The Length of the Word is ", length1)



					len1 := length1
					len2 := length2
					remaining := memory -(len1+len2)
					//totol length of the search 
					fmt.Println("First Word:", len1)
					fmt.Println("Second Word:", len2)
					fmt.Println("Memory Remaining:", remaining)

				}

			}

			//fmt.Println("Name", data)

		}
	}


	tmpl := template.Must(template.ParseFiles("templates/darknet.html"))

	tmpl.Execute(w, data)

}
