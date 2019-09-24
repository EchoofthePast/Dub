package handlers

import(
	"net/http"
	"fmt"
	"os"
	"html"
	"html/template"
	"strings"
	"github.com/Creator/DarkNet/static/goget"
)

//NewfileHandler New File Page Handler Function
func NewfileHandler(w http.ResponseWriter, r *http.Request) {

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

	var output [255]string
	var key string
	var value string
	var fname string
	var sname string
	var mname string
	var id string 
	var num string
	var age string
	var information string
	var medication string
	var tel string
	var cell string
	var address string
	var suburb string
	var city string
	var province string
	var country string


	if r.Method == "POST" {

		for k, v := range r.Form {
			key = k
			value = strings.Join(v, "")

			if key == "firstname" {
				output[0] = value
			} else if key == "surname" {
				output[1] = value
			} else if key == "middlename" {
				output[2] = value
			} else if key == "id" {
				output[3] = value
			} else if key == "number"{
				output[4] = value
			} else if key == "age" {
				output[5] = value
			} else if key == "tel" {
				output[6] = value
			} else if key == "cell" {
				output[7] = value
			} else if key == "address" {
				output[8] = value
			} else if key == "suburb" {
				output[9] = value
			} else if key == "city" {
				output[10] = value
			} else if key == "province" {
				output[11] = value
			} else if key == "country" {
				output[12] = value
			} else if key == "information" {
				output[13] = value
			} else if key == "medication" {
				output[14] = value
			}

			fmt.Println("key:", key)
			fmt.Println("value", value)
			//fmt.Println("val:", strings.Join(v, ""))
			fname = output[0]
			sname = output[1]
			mname = output[2]
			id = output[3]
			num = output[4]
			age = output[5]
			tel = output[6]
			cell = output[7]
			address = output[8]
			suburb = output[9]
			city = output[10]
			province = output[11]
			country = output[12]
			information = output[13]
			medication = output[14] 
			
		}

		fmt.Println(fname, sname, mname,id, num, age,tel, cell, address, suburb, city, province, country ,information, medication)
		goget.CreateFile(output)
	} //end of if statement

	tmpl := template.Must(template.ParseFiles("templates/newfile.html"))
	tmpl.Execute(w, nil)

}

