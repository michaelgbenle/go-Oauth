package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/michaelgbenle/go-Oauth/handler"
)

func SetupRouter() (*pat.Router,error){
	p := pat.New()
	p.Get("/auth/{provider}/callback", handler.SuccessHandler)

	p.Get("/auth/{provider}", handler.BeginAuth)

	p.Get("/", handler.HomePage)

return p, nil
}