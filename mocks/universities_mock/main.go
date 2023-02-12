package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func parseFile(filename string) ([]byte, error) {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return nil, e
	}
	return file, e
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		params := r.URL.Query()
		if len(params) == 0 || params["name"] == nil {

			http.Error(w, "No param or invalid param given this will return "+
				"all universities in the real API", http.StatusBadRequest)
			return
		}
		universityFilePath := "./res/university.json"

		if params["country"] != nil {
			universityFilePath = "./res/university_with_country.json"
		}

		log.Println("Received " + r.Method +
			" request on university mock, returning mock data.")

		w.Header().Add("content-type", "application/json")

		output, err := parseFile(universityFilePath)

		if err != nil {
			http.Error(w, "Error when reading resource file (Error: "+
				err.Error()+")", http.StatusInternalServerError)
			return
		}

		_, err2 := fmt.Fprint(w, string(output))
		if err2 != nil {
			http.Error(w, "Error when writing HTTP response (Error: "+
				err2.Error()+")", http.StatusInternalServerError)
			return

		}
	default:
		http.Error(w, "Handler not implemented", http.StatusNotImplemented)
	}
}

func main() {

	// Handle port assignment
	// (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	const DefaultPort = "8082"
	if port == "" {
		log.Println("$PORT has not been set. Default: " + DefaultPort)
		port = DefaultPort
	}

	// Set up handler endpoints
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/search/", searchHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
