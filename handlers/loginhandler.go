package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/Creator/DarkNet/static/goget"

)

//LoginHandler is the Handler for the login file
func LoginHandler(w http.ResponseWriter, r *http.Request){

	var data *goget.DataStructure
	var	flag *goget.Access
	var access string
	var status string
	var user string
	var pass string 
	
	r.ParseForm()
	
	//data = blank()

	//input from FrontEnd
	username := r.Form["username"] //

	///indexsearch(search)

	for index := range username {

		if username[index] != " " {

			input := username[index]

			user = input

			//data = goget.Gopher(input)
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

	password := r.Form["password"] //

	///indexsearch(search)

	for index := range password {

		if password[index] != " " {

			input := password[index]

			pass = input
			//data = goget.Gopher(input)
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


	//do a comparison between the data base users and the login username and password input
	fmt.Println("Username: ", user)
	fmt.Println("Password: ", pass)

	data = goget.Asroth("login")

	user0 := data.User0 
	password0 := data.Password0
	user1 := data.User1 
	password1 := data.Password1
	user2 := data.User2
	password2 := data.Password2
	

	if user0 == user{
		status, access = worker(user0,password0,user,pass)

	} else if user1 == user{
		status, access = worker(user1,password1,user, pass)

		} else if user2 == user{
			status, access = worker(user2,password2,user, pass)

			}


	flag = &goget.Access{
		Status: status,

	}
	fmt.Println("Gate: ", flag.Status)

	if access == "Granted"{
		http.HandleFunc("/index", IndexHandler)
		http.HandleFunc("/phonebookhandler", PhonebookHandler)
		http.HandleFunc("/darknet", DarknetHandler)
		http.HandleFunc("/newfile", NewfileHandler)
		http.HandleFunc("/input", InputHandler)
		http.HandleFunc("/upload", UploadHandler)
		http.HandleFunc("/edit", EditHandler)
		http.HandleFunc("/library", LibraryHandler)
		fmt.Println("Access: ", access)
	}else if access == "Denied"{
		fmt.Println("Access:", access)
		//page 404
	}
	
	
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, flag)
	
}

func worker(user0, password0, user, pass string)(result,access string){

	 fmt.Println("user0", user0)
	 fmt.Println("user", user)
	 fmt.Println("password0", password0)
	 fmt.Println("pass", pass)

	 //username and password comparison
	if user0 == user && password0 == pass{		
		//access = Flag.Status
		result = "Open"
		access = "Granted"
		}else{
			//access = Flag.Status
			result = "Closed"
			access = "Denied"
 	}

	return result,access
 
}