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
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	 googleId := os.Getenv("google-client-id")
  	 googleSecret := os.Getenv("google-client-secret")
	 port := ":" + os.Getenv("port")
	 key :=os.Getenv("Secret-session-key")  // Replace with your SESSION_SECRET or similar
  	 maxAge := 86400 * 30  // 30 days
  	 isProd := false       // Set to true when serving over https
	 store := sessions.NewCookieStore([]byte(key))
  	store.MaxAge(maxAge)
  	store.Options.Path = "/"
  	store.Options.HttpOnly = true   // HttpOnly should always be enabled
  	store.Options.Secure = isProd
  	gothic.Store = store

	callback_url := os.Getenv("callback-url")

  	goth.UseProviders(
    google.New(googleId, googleSecret, callback_url, "email", "profile"),
  	)
	p , err:=routes.SetupRouter()
	if err != nil {
		log.Fatal(err)
	}

 			log.Println("listening on localhost, port", port)
		log.Fatal(http.ListenAndServe(port, p))
}

