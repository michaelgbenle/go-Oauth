package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func SuccessHandler(res http.ResponseWriter, req *http.Request) {

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(res, user)
}

func BeginAuth(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}