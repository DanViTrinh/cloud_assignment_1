package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO: consider using get instead
func getResponseFromApi(w http.ResponseWriter,
	apiURL, apiName string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		http.Error(w, "Error in creating new request to "+apiName+" API",
			http.StatusInternalServerError)
		return nil
	}

	// request.Header.Add("content-type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	res, err := client.Do(request)

	if err != nil {
		http.Error(w, "Error in getting response from "+apiName+" API",
			http.StatusInternalServerError)
		return nil
	}
	return res
}

func populateDataWithResponse(w http.ResponseWriter, res *http.Response,
	uniName string, data interface{}) bool {

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, "Error during reading response",
			http.StatusInternalServerError)
		return false
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		http.Error(w, "Error during unmarshaling from "+uniName+" API",
			http.StatusInternalServerError)
		return false
	}
	return true
}

func UniInfoHandler(w http.ResponseWriter, r *http.Request) {
	uniUrl := "http://" + UniversitiesAPIurl + UniversitiesSearch
	uniName := "university"

	var unisFound []University

	res := getResponseFromApi(w, uniUrl, uniName)
	if res == nil {
		return
	}

	if !populateDataWithResponse(w, res, uniName, &unisFound) {
		return
	}
	// body, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	http.Error(w, "Error during reading response",
	// 		http.StatusInternalServerError)
	// 	return
	// }

	// err = json.Unmarshal(body, &unisFound)

	// if err != nil {
	// 	http.Error(w, "Error during unmarshaling from "+uniName+" API",
	// 		http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Add("content-type", "application/json")

	_, err := fmt.Fprintln(w, unisFound[0])
	if err != nil {
		http.Error(w, "Error during returning output",
			http.StatusInternalServerError)
		return
	}
}
