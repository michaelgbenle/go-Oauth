package main

import (
	"net/http"
	"log"
	"github.com/michaelgbenle/go-Oauth/routes"
)


func main() {

	p ,port, err:=routes.SetupRouter()
	if err != nil {
		log.Fatal(err)
	}

 		log.Println("listening on localhost port", port)
		log.Fatal(http.ListenAndServe(port, p))
}

