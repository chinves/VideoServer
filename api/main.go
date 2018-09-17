package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

		r := RegisterHandle()
		http.ListenAndServe(":28999",r)


}


// 路由
func RegisterHandle() *httprouter.Router{

	r := httprouter.New()
	r.POST("/",CreateUseruu)
	return r
}


