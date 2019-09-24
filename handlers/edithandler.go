package handlers

import(
	
	"os"
	"fmt"
	"html"
	"strings"
	"net/http"
	"html/template"
	"github.com/Creator/Dub/static/goget"
	
)
//EditHandler Edit Page Handler Function
func EditHandler(w http.ResponseWriter, r *http.Request) {
	
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
						letters[index] = string(input[index])
						
						var i int
						var length1 int
						var length2 int
						var words string

						if letters[index] == " " {
							
							fmt.Println("Search Word ", index, letters[index])
							i++
							length1 = index
							store := length1
							//fmt.Println("The Length of the Word is ", length1)
							
							if index == store{
								
								if letters[index] != " " {
									for index = range letters{
										words = letters[index]
									}
									
									fmt.Println("Search Word ", index, words)
									i++
									length2 = index
									//fmt.Println("The Length of the Word is ", length2)
									
									
								}
								
							}
							
							
							len1 := length1
							len2 := length2
							memory := len(letters)
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
			
			var output [255]string
			var key string
			var value string
			var sname string
			var fname string
			var mname string
			var id string
			var num string
			var age string
			var tel string
			var cell string
			var address string
			var suburb string
			var city string
			var province string
			var country string
			var information string
			var medication string
			
			//var add0 string
			//var add1 string
			//var add2 string
			var add3 string
			var add4 string
			var doc0 string
			var doc1 string
			var doc2 string
			
			
			//get the input from the button arrays
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
					} else if key == "number" {
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
																} else if key == "doc0" {
																	output[15] = value
																	} else if key == "doc1" {
																		output[16] = value
																		} else if key == "doc2" {
																			output[17] = value
																			} else if key == "doc3" {
																				output[18] = value
																			}
						fmt.Println("key:", key)
						fmt.Println("value", value)

						fname = output[0]
						sname = output[1]
						mname = output[2]
						id = output[3]
						num = output[4]
						age = output[5]
						tel = output[6]	
						cell = output[7]
						address = output[8]
						suburb =  output[9]
						city = output[10]
						province = output[11]
						country = output[12]
						information = output[13]
						medication = output[14]
						
						doc0 = output[15]
						doc1 = output[16]
						doc2 = output[17]
					}
					
					fmt.Println(fname, sname, mname,id, num, age,tel, cell,address,suburb, city, province, country ,information, medication,add3, add4, doc0, doc1,doc2)
					goget.CreateFile(output)
				} //end of if statement
					
					tmpl := template.Must(template.ParseFiles("templates/edit.html"))
					tmpl.Execute(w, data)
					
}//EditHandler Function End
				