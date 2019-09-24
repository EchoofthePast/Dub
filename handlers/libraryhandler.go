package handlers

import(
	"net/http"
	"fmt"
	"os"
	"html"
	"html/template"
	"strings"
	"github.com/Creator/Dub/static/goget"
)
var data *goget.DataStruct

//LibraryHandler Input Page Handler Function
func LibraryHandler(w http.ResponseWriter, r *http.Request) {
	
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



	r.ParseForm()

	if r.Method == "GET" {
		fmt.Fprintf(os.Stdout, "GET, %q", html.EscapeString(r.URL.Path))
	} else if r.Method == "POST" {
		fmt.Fprintf(os.Stdout, "POST, %q", html.EscapeString(r.URL.Path))
	} else {
		http.Error(w, "Invalid request method.", 405)
	}

	fmt.Println(" Method:", method)

	
	var (
		output [255]string
		key string
		value string
		oName string
		oType string
		oClass string
		oPurpose string 
		oDescription string
		oAge string
		oComponents string
		oMaterials string
		oInformation string
		oDimensions string
		oFound string
		oCreated string
		// add0 string
		// add1 string
		// add2 string

	)	

	


	if r.Method == "POST" {

		for k, v := range r.Form {
			key = k
			value = strings.Join(v, "")

			if key == "name" {
				output[0] = value
			} else if key == "type" {
				output[1] = value
			} else if key == "class" {
				output[2] = value
			} else if key == "purpose" {
				output[3] = value
			} else if key == "description"{
				output[4] = value
			} else if key == "age" {
				output[5] = value
			} else if key == "components" {
				output[6] = value
			} else if key == "materials" {
				output[7] = value
			} else if key == "information" {
				output[8] = value
			} else if key == "dimensions" {
				output[9] = value
			} else if key == "found" {
				output[10] = value
			} else if key == "created" {
				output[11] = value
			//} else if key == "add0" {
			//	output[12] = value
			//} else if key == "add1" {
			//	output[13] = value
			//} else if key == "add2" {
			//	output[14] = value
			}

			fmt.Println("key:", key)
			fmt.Println("value", value)
			//fmt.Println("val:", strings.Join(v, ""))
			oName = output[0]
			oType = output[1]
			oClass = output[2]
			oPurpose = output[3]
			oDescription = output[4]
			oAge = output[5]
			oComponents = output[6]
			oMaterials = output[7]
			oInformation = output[8]
			oDimensions = output[9]
			oFound = output[10]
			oCreated = output[11]
			//add0 = output[12]
			//add1 = output[13]
			//add2 = output[14] 
			
		}

		fmt.Println(oName,oType,oClass,oPurpose,oDescription,oAge,oComponents,oMaterials,oInformation,oDimensions,oFound,oCreated)
		goget.ObjectFile(output)
	} //end of if statement

	tmpl := template.Must(template.ParseFiles("templates/library.html"))
	tmpl.Execute(w, nil)

}

