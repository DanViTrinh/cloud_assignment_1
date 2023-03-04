package main

import (
	"log"
	"net/http"
	"os"
	h "university_service/handlers"
	u "university_service/handlers/utilities"
)

func main() {
	// Retrieve the potential environment variable
	port := os.Getenv("PORT")
	defaultPort := "8083"
	if port == "" {
		log.Println("$PORT has not been set. Default:" + defaultPort)
		port = defaultPort
	}

	// Handler endpoints
	http.Handle(u.DefaultPath, h.RootHandler(h.EmptyHandler))

	http.Handle(u.UniInfoPath, h.RootHandler(h.UniInfoHandler))
	http.Handle(u.UniInfoPath+"/", h.RootHandler(h.UniInfoHandler))

	http.Handle(u.NeighborUnisPath, h.RootHandler(h.NeighborUniHandler))
	http.Handle(u.NeighborUnisPath+"/", h.RootHandler(h.NeighborUniHandler))

	http.Handle(u.DiagPath, h.RootHandler(h.DiagHandler))
	http.Handle(u.DiagPath+"/", h.RootHandler(h.DiagHandler))

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
