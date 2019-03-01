package main

import (
	"io"
	"log"
	"net/http"
)

func HttpHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"hello world")
}

func main(){
	http.HandleFunc("/hello",HttpHandler)
	err  := http.ListenAndServeTLS(":9010","/Users/zhan/cert2/server.crt","/Users/zhan/cert2/server.key",nil)
	if err != nil{
		log.Fatal("ListenAndServer : ",err.Error())
	}
}
