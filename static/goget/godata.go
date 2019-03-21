package goget

import (
	"bufio"
	"fmt"
	"os"
)

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
