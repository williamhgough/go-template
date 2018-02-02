package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is index router.")
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Printf("Running api server in %s mode", env)
	}

	http.ListenAndServe(":8080", router)
}
