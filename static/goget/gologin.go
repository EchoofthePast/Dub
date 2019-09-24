package goget

import(
	"os"
	"bufio"
	"fmt"

)

//Asroth Searches for file and then converts the data and returns
func Asroth(input string) (data *DataStructure) {

	//file := fileSearcher(input)

	var src [255]string
	var out [255]string
	//fmt.Println("File ", file)


	//search for file by input or varias combinations
	path := "static/goget/scrap/" + input + ".txt"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Sorry I could not locate that file for you", err)
	} else if err == nil {

		fmt.Println("File located")
	}
		var lines []string
	//var stringline []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//data is passing from file to lines
	for index := range lines {
		out[index] = lines[index]

		//fmt.Println(data)
	}


	for index := range out {
		src[index] = out[index]
	}

	user0 := src[0]
	password0 := src[1]
	user1 := src[2]
	password1 := src[3]
	user2 := src[4]
	password2 := src[5]
	user3 := src[6]
	password3 := src[7]
	user4 := src[8]
	password4 := src[9]
	user5 := src[10]
	password5 := src[11]
	user6 := src[12]
	password6 := src[13]
	user7 := src[14]
	password7 := src[15]
	user8 := src[16]
	password8 := src[17]
	user9 := src[18]
	password9 := src[19]


	data = &DataStructure{

			User0: user0,
			Password0: password0,
			User1: user1,
			Password1: password1,
			User2: user2,
			Password2: password2,
			User3: user3,
			Password3: password3,
			User4: user4,
			Password4: password4,
			User5: user5,
			Password5: password5,
			User6: user6,
			Password6: password6,
			User7: user7,
			Password7: password7,
			User8: user8,
			Password8: password8,
			User9: user9,
			Password9: password9,
	}

	return data

}

