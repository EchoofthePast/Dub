package goget

import (
	"fmt"
	"bufio"
	"os"

)

//CreateFile a new data file and writes the data to it
func CreateFile(output [255]string) {

	//Create a new file
	fname := output[0]
	sname := output[1]
	mname := output[2]
	id := output[3]
	num := output[4]
	age := output[5]
	tel := output[6] 
	cell := output[7]
	address := output[8]
	suburb := output[9]
	city := output[10]
	province := output[11]
	country := output[12]
	info := output[13]
	meds := output[14]
	add3 := output[15]
	add4 := output[16]
	doc0 := output[17]
	doc1 := output[18]
	doc2 := output[19]
	// A  little bit of Coding magic and we have a file creator 	
	path := "static/goget/files/"+ fname + " " + sname  + ".txt"
	
	file, err := os.Create(path)
	if err != nil{
		fmt.Println("Error Creating File", err)
	}
	
	//Create a new writer.
	w := bufio.NewWriter(file)
	
	// Write a string to the file.
	w.WriteString("static/images/"+ fname + " " + sname  +  ".jpg\n")
	w.WriteString(fname + "\n")
	w.WriteString(sname + "\n")
	w.WriteString(mname + "\n")
	w.WriteString(id + "\n")
	w.WriteString(num + "\n")
	w.WriteString(age + "\n")
	w.WriteString(tel + "\n")
	w.WriteString(cell + "\n")
	w.WriteString(address + "\n")
	w.WriteString(suburb + "\n")
	w.WriteString(city + "\n")
	w.WriteString(province + "\n")
	w.WriteString(country + "\n")
	w.WriteString(info + "\n")
	w.WriteString(meds + "\n")
	w.WriteString(add3 + "\n")
	w.WriteString(add4 + "\n")
	w.WriteString(doc0 + "\n")
	w.WriteString(doc1 + "\n")
	w.WriteString(doc2 + "\n")
	
	// Flush.
	w.Flush()
}


//ObjectFile a new data file and writes the data to it
func ObjectFile(output [255]string) {

	//Create a new file
	fname := output[0]
	sname := output[1]
	mname := output[2]
	id := output[3]
	num := output[4]
	age := output[5]
	tel := output[6] 
	cell := output[7]
	address := output[8]
	suburb := output[9]
	city := output[10]
	province := output[11]
	country := output[12]
	info := output[13]
	meds := output[14]
	add3 := output[15]
	add4 := output[16]
	doc0 := output[17]
	doc1 := output[18]
	doc2 := output[19]
	// A  little bit of Coding magic and we have a file creator 	
	path := "static/goget/object files/"+ fname + " " + sname  + ".txt"
	
	file, err := os.Create(path)
	if err != nil{
		fmt.Println("Error Creating File", err)
	}
	
	//Create a new writer.
	w := bufio.NewWriter(file)
	
	// Write a string to the file.
	w.WriteString("static/images/"+ fname + " " + sname  +  ".jpg\n")
	w.WriteString(fname + "\n")
	w.WriteString(sname + "\n")
	w.WriteString(mname + "\n")
	w.WriteString(id + "\n")
	w.WriteString(num + "\n")
	w.WriteString(age + "\n")
	w.WriteString(tel + "\n")
	w.WriteString(cell + "\n")
	w.WriteString(address + "\n")
	w.WriteString(suburb + "\n")
	w.WriteString(city + "\n")
	w.WriteString(province + "\n")
	w.WriteString(country + "\n")
	w.WriteString(info + "\n")
	w.WriteString(meds + "\n")
	w.WriteString(add3 + "\n")
	w.WriteString(add4 + "\n")
	w.WriteString(doc0 + "\n")
	w.WriteString(doc1 + "\n")
	w.WriteString(doc2 + "\n")
	
	// Flush.
	w.Flush()
}
