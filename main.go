package main

import (
	"assignment1/router"
	"log"
	"net/http"
	"os"
)

func main() {
	//will listen on port 4000
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Fatal(http.ListenAndServe(":"+port, router.Router()))
}
