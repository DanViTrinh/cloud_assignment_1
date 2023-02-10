package main

import (
	"log"
	"net/http"
	"os"
	"university_service/handlers"
)

func main() {
	// Retrieve the potential enviroment variable
	port := os.Getenv("PORT")
	defaultPort := "8080"
	if port == "" {
		log.Println("$PORT has not been set. Default:" + defaultPort)
		port = defaultPort
	}

	// Handler endpoints
	http.HandleFunc(handlers.DefaultPath, handlers.EmptyHandler)

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
