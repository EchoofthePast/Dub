package goget



type DataStruct struct {
	Name    string
	Surname string
	Image1  string
	Image2  string
	Image3  string
	Power   string
	Gold    string
	Men     string
	Both    string
	Url     string
	Info    string
	Doc     string
	Life    string

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
	doc := src[5]
	life := src[6]
	power := src[7]
	gold := src[8]
	men := src[9]

	data = &DataStruct{

		Name:    name,
		Surname: surname,
		Both:    name + surname,
		Url:     "",
		Image1:  image1,
		Image2:  image2,
		Info:    info,
		Doc:     doc,
		Life:    life,
		Power:   power,
		Gold:    gold,
		Men:     men,

	}

	return data

}
