package handlers

import(
	"net/http"
	"fmt"
	"os"
	"io"
	"html/template"
)



//UploadHandler Upload Page Handler Function
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		//crutime := time.Now().Unix()
		//h := md5.New()
		//io.WriteString(h, strconv.FormatInt(crutime, 10))
		//token := fmt.Sprintf("%x", h.Sum(nil))

		tmpl := template.Must(template.ParseFiles("templates/upload.html"))
		tmpl.Execute(w, data)

	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
