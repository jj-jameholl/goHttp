 package main

 import (
	 "log"
	 "net"
	 "net/http"
	 "net/rpc"
 )

type Arg struct {
	A,B int
}

type Reply struct {
	Code int
	Data int
}

type Server int

func (server *Server) Add(args *Arg , reply *Reply) error {
	reply.Data = args.A + args.B
	reply.Code = 200
	return nil
}

func (server *Server ) Multi(args *Arg , reply *Reply) error {
	reply.Data = args.A * args.B
	reply.Code = 200
	return nil
}

func main(){
	ser := new(Server)
	rpc.Register(ser)
	rpc.HandleHTTP()
	l , e := net.Listen("tcp" ,":1234")

	if e != nil {
		log.Fatal("listen failed :", e)
	}
	go http.Serve(l,nil)

	i := 10
	for i>0 {
		i ++
	}
}


