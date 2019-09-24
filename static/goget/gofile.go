package goget

import (
	"bufio"
	"fmt"
	"os"
)

//Gopher Reads the File with a pointer to the Global DataStruct.
func Gopher(input string) (data *DataStruct) {

	file := fileSearcher(input)

	var src [255]string

	for index := range file {
		src[index] = file[index]
	}

	image := src[0]
	name := src[1]
	surname := src[2]
	middlename := src[3]
	id := src[4]
	number := src[5]
	age := src[6]
	tel := src[7]
	cell := src[8]
	address :=src[9]
	suburb := src[10]
	city := src[11]
	province := src[12]
	country := src[13]
	information := src[14]
	medication := src[15]
	add3 := src[16]
	add4 := src[17]
	doc0 := src[18]
	doc1 := src[19]
	doc2 := src[20]


	data = &DataStruct{

		Image:      image,
		Name:       name,
		Surname:    surname,
		Both:       name + " " + surname,
		Middlename: middlename,
		ID:			id,
		Number:     number,
		Age:		age,
		Tel:  		tel,
		Cell:		cell,
		Address:    address,
		Suburb: 	suburb,
		City:		city,
		Province:   province,
		Country:    country,
		Information:information,
		Medication: medication,
		Add3:       add3,
		Add4:		add4,
		Doc0:		doc0,
		Doc1: 		doc1,
		Doc2:		doc2,

	}

	return data
}

func fileSearcher(input string) (info [255]string) {

	info = Filer(input)
	return info

}

//Filer search data export
func Filer(input string) (data [255]string) {

	file := search(input)
	//					//
	data = readFromFile(file)
	//					//
	//fmt.Println("Data filer:", data)

	return data
}

func search(input string) (file *os.File) {
	//search for file by input or varias combinations
	fmt.Println("Search Running..")
	path := "static/goget/files/" + input + ".txt"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Sorry I could not locate that file for you", err)
	} else if err == nil {

		fmt.Println("File located")
	}

	//defer file.Close()
	return file
}

func readFromFile(file *os.File) (data [255]string) {

	var lines []string
	//var stringline []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//data is passing from file to lines
	for index := range lines {
		data[index] = lines[index]

		//fmt.Println(data)
	}
	return data
}
