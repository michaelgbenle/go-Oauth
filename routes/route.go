package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/markbates/goth/gothic"
	"github.com/michaelgbenle/go-Oauth/handler"
)

func SetupRouter() {
	p := pat.New()
	p.Get("/auth/{provider}/callback", handler.SuccessHandler)

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(res, false)
	})

}