package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"net"
)

func main(){
	if len(os.Args) != 2{
		fmt.Println("it is wrong")
		os.Exit(1)
	}

	service := os.Args[1]

	conn,err := net.Dial("tcp",service)
	checkErr(err)

	_,err1 := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err1)

	resp,err := fullRead(conn)
	checkErr(err)

	fmt.Println(resp)
}

func fullRead(conn net.Conn) (resp string , err error){
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n,err := conn.Read(buf[0:])


		if err != nil {
			if err == io.EOF{
				break
			}
		}

		result.Write(buf[0:n])
	}
	b := result.Bytes()
	resp = string(b)
	return
}

func checkErr(err error) {
	if err != nil{
		if err == io.EOF{
			fmt.Printf("something is wrong:%s",err.Error())
		}
		os.Exit(1)
	}
}
