package goget

import (
	"bufio"
	"os"
)

//SmallData player selection data struct
type SmallData struct {
	Selection0 string
	Selection1 string
	Selection2 string
	Selection3 string
	Selection4 string
}

//CreateFile opens a datafile and writes to it
func CreateFile(input string, player1 *SmallData) {

	player1.Selection0 = ""
	player1.Selection1 = ""
	player1.Selection2 = ""

	path := "static/goget/files/" + input + ".txt"

	//Create a new file
	f, _ := os.Create(path)

	//Create a new writer.
	w := bufio.NewWriter(f)

	// Write a string to the file.
	w.WriteString(player1.Selection0)
	w.WriteString(player1.Selection1)
	w.WriteString(player1.Selection2)

	// Flush.
	w.Flush()
}
