package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func ParseFile(filename string) ([]byte, error) {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return nil, e
	}
	return file, nil
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
		output, err := ParseFile("./res/norway.json")

		if err != nil {
			http.Error(w, "Error when reading resource file (Error: "+
				err.Error()+")", http.StatusInternalServerError)
			return
		}

		_, err2 := fmt.Fprint(w, string(output))
		if err2 != nil {
			http.Error(w, "Error when writing HTTP response (Error: "+
				err.Error()+")", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Handler not implemented", http.StatusNotImplemented)
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

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
