package main

import (
	"log"
	"net/http"
	"os"
	"university_service/handlers"
	"university_service/handlers/utilities"
)

func main() {
	// Retrieve the potential enviroment variable
	port := os.Getenv("PORT")
	defaultPort := "8083"
	if port == "" {
		log.Println("$PORT has not been set. Default:" + defaultPort)
		port = defaultPort
	}

	// Handler endpoints
	http.HandleFunc(utilities.DefaultPath, handlers.EmptyHandler)

	http.HandleFunc(utilities.UniInfoPath, handlers.UniInfoHandler)
	//TODO: TEST if it works without the /
	http.HandleFunc(utilities.UniInfoPath+"/", handlers.UniInfoHandler)

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
