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
const SeveralCountryFilePath string = "./res/united_countries.json"
const HandlerNotImplementedMessage string = "Handler not implemented"

func parseFile(filename string) ([]byte, error) {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return nil, e
	}
	return file, nil
}

func readErrorHandler(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, "Error when reading resource file (Error: "+
			err.Error()+")", http.StatusInternalServerError)
		return false
	}
	return true
}

func writeHttpResponseErrorHandler(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, "Error when writing HTTP response (Error: "+
			err.Error()+")", http.StatusInternalServerError)
		return false
	}
	return true
}

func searchName(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + r.Method)
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 4 || parts[3] == "" {
			status := http.StatusBadRequest
			http.Error(w, "Expecting .../name, cannot be empty", status)
			return
		}

		w.Header().Add("content-type", "application/json")
		var outputFile string

		fullTextParam := r.URL.Query()["fullText"]
		if len(r.URL.RawQuery) != 0 && fullTextParam != nil &&
			fullTextParam[0] == "true" {
			outputFile = SingleCountryFilePath
		} else {
			outputFile = SeveralCountryFilePath
		}

		output, err := parseFile(outputFile)

		if !readErrorHandler(err, w) {
			return
		}

		_, err2 := fmt.Fprint(w, string(output))
		if !writeHttpResponseErrorHandler(err2, w) {
			return
		}
	default:
		http.Error(w, HandlerNotImplementedMessage, http.StatusNotImplemented)
	}
}

func searchByCode(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + r.Method)
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 4 || parts[3] == "" {
			status := http.StatusBadRequest
			http.Error(w, "Expecting .../code, cannot be empty", status)
			return
		}
		w.Header().Add("content-type", "application/json")

		output, err := parseFile(SingleCountryFilePath)
		if !readErrorHandler(err, w) {
			return
		}
		_, err2 := fmt.Fprint(w, string(output))

		if !writeHttpResponseErrorHandler(err2, w) {
			return
		}
	default:
		http.Error(w, HandlerNotImplementedMessage, http.StatusNotImplemented)
	}

}

func main() {
	port := os.Getenv("PORT")
	const DefaultPort = "8081"
	if port == "" {
		log.Println("$PORT has not been set. Default: " + DefaultPort)
		port = DefaultPort
	}

	http.HandleFunc("/v3.1/name", searchName)
	http.HandleFunc("/v3.1/name/", searchName)

	http.HandleFunc("/v3.1/alpha", searchByCode)
	http.HandleFunc("/v3.1/alpha/", searchByCode)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
