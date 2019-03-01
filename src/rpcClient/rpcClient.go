package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Reply struct {
	Code int
	Data int
}

type Arg struct {
	A,B int
}

func main(){
	client ,err := rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing :",err)
	}

	args := &Arg{7,8}
	reply := &Reply{}

	err = client.Call("Server.Multi",args,reply)

	if err != nil {
		log.Fatal("server err :",err)
	}

	fmt.Printf("return code is %d,and data is %d",reply.Code,reply.Data)
}