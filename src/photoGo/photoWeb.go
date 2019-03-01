package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)
func uploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		w.Header().Set("Content-Type","text/html;charset=utf-8")
		io.WriteString(w,
			"<form method=\"POST\" action=\"/uploads\" enctype=\"multipart/form-data\">" +
				"Choose an image to uploads: <input name=\"image\" type=\"file\" />" +
				"<input type=\"submit\" value=\"Upload\" />" +
				"</form>")
		return
	}

	if r.Method == "POST" {
		f,h,err := r.FormFile("image")
		if err != nil {
			http.Error(w,err.Error(),
			http.StatusInternalServerError)

			return
		}

		filename := h.Filename
		defer f.Close()
		t ,err := os.Create(UPLOAD_DIR + "/" + filename)
		if t != nil {
			http.Error(w,err.Error(),
			http.StatusInternalServerError)

			return
		}

		if _,err := io.Copy(t,f);err != nil {
			http.Error(w,err.Error(),
			http.StatusInternalServerError)

			return
		}

		http.Redirect(w,r,"/view?id="+filename, http.StatusFound)
	}
}


func viewHandler(w http.ResponseWriter,r *http.Request) {
	imageId := r.FormValue("id")
	imagePath :=  UPLOAD_DIR + "/" + imageId

	if exist := isExist(imagePath);!exist{
		http.NotFound(w,r)
		return
	}
	w.Header().Set("Content-Type","image")
	http.ServeFile(w,r,imagePath)
	return
}


func listHandler(w http.ResponseWriter,r *http.Request) {
	fileInfoArr , err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	var listHtml string
	for _,fileInfo := range fileInfoArr{
		imageid := fileInfo.Name()
		listHtml += "<li><a href = \"/view?id="+ imageid +"\">imageId</a></li>"
	}

	w.Header().Set("Content-Type","text/html;charset=utf-8")
	io.WriteString(w,"<ol>" + listHtml + "</ol>")
}


func isExist(path string) bool{
	_,err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)

}



func main(){
	http.HandleFunc("/uploads",uploadHandler)
	http.HandleFunc("/view",viewHandler)
	http.HandleFunc("/list",listHandler)

	err := http.ListenAndServe(":8010",nil)

	if err != nil {
		log.Fatal("err happep :",err.Error())
	}
}