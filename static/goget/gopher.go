package goget



type DataStruct struct {
	Name    string
	Surname string
	Image1  string
	Image2  string
	Image3  string
	Both    string
	ID     string
	Info    string


}

func fileSearcher(input string) (info [255]string) {

	info = Filer(input)
	//fmt.Println("Information", info)
	return info

}

func Gopher(input string) (data *DataStruct) {

	file := fileSearcher(input)

	var src [255]string
	//fmt.Println("File ", file)

	for index := range file {
		src[index] = file[index]
	}

	image1 := src[0]
	image2 := src[1]
	name := src[2]
	surname := src[3]
	info := src[4]
	id := src[5]

	data = &DataStruct{

		Name:    name,
		Surname: surname,
		Both:    name + surname,
		Image1:  image1,
		Image2:  image2,
		Info:    info,
		ID:		 id,

	}

	return data

}
