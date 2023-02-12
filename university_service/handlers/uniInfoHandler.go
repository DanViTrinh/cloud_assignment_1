package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UniInfoHandler(w http.ResponseWriter, r *http.Request) {
	uniUrl := "http://" + UniversitiesAPIurl + UniversitiesSearch

	//TODO: consider using get instead
	// Create new request
	uniRequest, err := http.NewRequest(http.MethodGet,
		uniUrl, nil)
	if err != nil {
		http.Error(w, "Error in creating new request to universities API",
			http.StatusInternalServerError)
		return
	}

	uniRequest.Header.Add("content-type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	res, err := client.Do(uniRequest)

	if err != nil {
		http.Error(w, "Error in getting response from universities API",
			http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error during reading response",
			http.StatusInternalServerError)
		return
	}

	var unisFound []University

	err = json.Unmarshal(body, &unisFound)

	if err != nil {
		http.Error(w, "Error during unmarshaling from universities API",
			http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")

	_, err = fmt.Fprintln(w, unisFound[0])
	if err != nil {
		http.Error(w, "Error during returning output",
			http.StatusInternalServerError)
		return
	}
}
