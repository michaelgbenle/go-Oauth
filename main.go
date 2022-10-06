package main

import (
	_ "fmt"
	_ "html/template"
	"net/http"
	
	"os"

	"log"

	_ "github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/michaelgbenle/go-Oauth/routes"
)


func main() {

	p ,port, err:=routes.SetupRouter()
	if err != nil {
		log.Fatal(err)
	}

 		log.Println("listening on localhost, port", port)
		log.Fatal(http.ListenAndServe(port, p))
}

