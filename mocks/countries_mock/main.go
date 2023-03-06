package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const SingleCountryFilePath string = "./res/norway.json"
const SeveralCountryFilePath string = "./res/s_countries.json"
const HandlerNotImplementedMessage string = "Handler not implemented"
const RootPath string = "/v3.1/"

// Parsing a json file
func parseFile(filename string) ([]byte, error) {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return nil, e
	}
	return file, nil
}

// Read error
func readErrorHandler(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, "Error when reading resource file (Error: "+
			err.Error()+")", http.StatusInternalServerError)
		return false
	}
	return true
}

// Writes a http error response
func writeHttpResponseErrorHandler(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, "Error when writing HTTP response (Error: "+
			err.Error()+")", http.StatusInternalServerError)
		return false
	}
	return true
}

// Displays a json file
func displayFileHandler(w http.ResponseWriter, filepath string) {
	w.Header().Add("content-type", "application/json")
	output, err := parseFile(filepath)

	if !readErrorHandler(err, w) {
		return
	}
	_, err2 := fmt.Fprint(w, string(output))
	if !writeHttpResponseErrorHandler(err2, w) {
		return
	}
}

// Checks whether the parameters in the url is valid
func checkParam(w http.ResponseWriter, urlPath string,
	desiredLen int, expecting string) bool {

	parts := strings.Split(urlPath, "/")
	if len(parts) != desiredLen || parts[desiredLen-1] == "" {
		status := http.StatusBadRequest
		http.Error(w, "Expecting .../"+expecting+
			", cannot be empty", status)
		return false
	}
	if parts[desiredLen-1] == "empty" {
		http.Error(w, "Not found", http.StatusNotFound)
		return false
	}
	return true
}

// Search by name displays one country, if the name of that country is empty
// status not found will be displayed
func searchName(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + r.Method)

		if !checkParam(w, r.URL.Path, 4, "name") {
			return
		}

		var outputFile string

		fullTextParam := r.URL.Query()["fullText"]
		if len(r.URL.Query()) != 0 && fullTextParam != nil &&
			fullTextParam[0] == "true" {
			outputFile = SingleCountryFilePath
		} else {
			outputFile = SeveralCountryFilePath
		}

		displayFileHandler(w, outputFile)

	default:
		http.Error(w, HandlerNotImplementedMessage, http.StatusNotImplemented)
	}
}

// search by code displays on one single country
func searchByCode(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + r.Method)
		if !checkParam(w, r.URL.Path, 4, "code") {
			return
		}

		displayFileHandler(w, SingleCountryFilePath)
	default:
		http.Error(w, HandlerNotImplementedMessage, http.StatusNotImplemented)
	}

}

// handles search by sub region and returns several countries
func searchBySubregion(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + r.Method + " on subregion")
		if !checkParam(w, r.URL.Path, 4, "subregion") {
			return
		}

		displayFileHandler(w, SeveralCountryFilePath)
	}
}

func main() {
	port := os.Getenv("PORT")
	const DefaultPort = "8081"
	if port == "" {
		log.Println("$PORT has not been set. Default: " + DefaultPort)
		port = DefaultPort
	}

	http.HandleFunc(RootPath+"name", searchName)
	http.HandleFunc(RootPath+"name/", searchName)

	http.HandleFunc(RootPath+"alpha", searchByCode)
	http.HandleFunc(RootPath+"alpha/", searchByCode)

	http.HandleFunc(RootPath+"subregion", searchBySubregion)
	http.HandleFunc(RootPath+"subregion/", searchBySubregion)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
