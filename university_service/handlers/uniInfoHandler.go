package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

// TODO: consider using get instead
func getResponseFromApi(w http.ResponseWriter,
	apiURL, apiName string, params map[string]string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		http.Error(w, "Error in creating new request to "+apiName+" API",
			http.StatusInternalServerError)
		return nil
	}

	q := request.URL.Query()

	for key, val := range params {
		q.Add(key, val)
	}

	request.URL.RawQuery = q.Encode()

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

func handleGetUniInfo(w http.ResponseWriter, r *http.Request) {
	uniUrl := "http://" + UniversitiesAPIurl + UniversitiesSearch
	uniName := "university"

	//TODO: get params from user url
	dummyParams := make(map[string]string)
	dummyParams["name"] = "random"
	dummyParams["country"] = "some country"

	res := getResponseFromApi(w, uniUrl, uniName, dummyParams)
	if res == nil {
		return
	}

	var unisFound []University

	if !populateDataWithResponse(w, res, uniName, &unisFound) {
		return
	}

	for _, uniFound := range unisFound {
		if !marshalAndDisplayData(w, uniFound) {
			return
		}
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
