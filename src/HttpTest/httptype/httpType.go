package httptype

import (
	"net/http"
)


func Send(addr string,hm string,params string) (r *http.Response,err error){
	switch hm {
	case "POST":
		r,err = postsend(addr,params)
	case "GET":
		r,err = getsend(addr)
	}
	return
}


func postsend(addr,params string) (r *http.Response,err error){
	return
}

func getsend(addr string) (r *http.Response,err error){
	r,err = http.Get(addr)
	if err != nil{
		return
	}
	return
}

