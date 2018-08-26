package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	io.WriteString(w, "create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	userName := p.ByName("user_name")

	io.WriteString(w, userName)
}

