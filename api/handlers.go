package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateUseruu(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	sendNormalResponse(w, string("create user"), 201)
}
