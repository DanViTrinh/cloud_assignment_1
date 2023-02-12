package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func getParamFromRequestURL(r *http.Request, desiredLen int) string {
	parts := strings.Split(r.URL.Path, "/")

	if (len(parts) == desiredLen && parts[desiredLen-1] != "") ||
		(len(parts) == desiredLen+1 && parts[desiredLen-1] != "" &&
			parts[desiredLen] == "") {
		return parts[desiredLen-1]
	}
	return ""
}

// TODO: optional fix: Sometimes getting duplicate universities from real api
func handleGetUniInfo(w http.ResponseWriter, r *http.Request) {
	uniUrl := UniversitiesAPIurl + UniversitiesSearch
	uniName := "university"

	name := getParamFromRequestURL(r, 5)
	if name == "" {
		http.Error(w, "Expecting format .../{university_name}",
			http.StatusBadRequest)
		return
	}

	params := make(map[string]string)
	params["name"] = name

	res := getResponseFromApi(w, uniUrl, uniName, params)
	if res == nil {
		return
	}

	var unisFound []University

	if !populateDataWithResponse(w, res, uniName, &unisFound) {
		return
	}

	if !marshalAndDisplayData(w, unisFound) {
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
