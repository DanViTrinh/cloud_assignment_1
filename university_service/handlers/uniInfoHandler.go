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

func marshalAndDisplayData(w http.ResponseWriter, data interface{}) bool {
	w.Header().Add("content-type", "application/json")

	jsonEncodedData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Error during marshalling data",
			http.StatusInternalServerError)
		return false
	}

	_, err = fmt.Fprint(w, string(jsonEncodedData))
	if err != nil {
		http.Error(w, "Error during writing response",
			http.StatusInternalServerError)
		return false
	}
	return true
}

func handleGetUniInfo(w http.ResponseWriter, r *http.Request) {
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

	if !marshalAndDisplayData(w, unisFound[0]) {
		return
	}
}
func UniInfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetUniInfo(w, r)
	default:
		http.Error(w, "Method not yet supported ", http.StatusNotImplemented)
	}
}
