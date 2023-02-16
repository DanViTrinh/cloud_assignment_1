package main

import (
	"log"
	"net/http"
	"os"
	"university_service/handlers"
	"university_service/handlers/utilities"
)

//TODO: handle errors better
// https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7?fbclid=IwAR1pa3LUFcoRvZoZ8kuGRcrRfTFf_5xWZRqE1Vy9DskYw2MOc9vl_JPWd7Y
// https://medium.com/@ozdemir.zynl/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd

func main() {
	// Retrieve the potential enviroment variable
	port := os.Getenv("PORT")
	defaultPort := "8083"
	if port == "" {
		log.Println("$PORT has not been set. Default:" + defaultPort)
		port = defaultPort
	}

	// Handler endpoints
	http.Handle(utilities.DefaultPath, handlers.RootHandler(handlers.EmptyHandler))

	http.Handle(utilities.UniInfoPath, handlers.RootHandler(handlers.UniInfoHandler))
	//TODO: TEST if it works without the /
	http.Handle(utilities.UniInfoPath+"/", handlers.RootHandler(handlers.UniInfoHandler))

	http.Handle(utilities.NeighbourUnisPath,
		handlers.RootHandler(handlers.NeighbourUniHandler))
	http.Handle(utilities.NeighbourUnisPath+"/",
		handlers.RootHandler(handlers.NeighbourUniHandler))

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
