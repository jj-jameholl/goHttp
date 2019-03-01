package main

import (
	"io"
	"os"
	"HttpTest/httptype"
		)

func main(){
	addr := os.Args[1]
	hm := os.Args[2]
	params := os.Args[3]
	resp,_ := httptype.Send(addr,hm,params)
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)
}

