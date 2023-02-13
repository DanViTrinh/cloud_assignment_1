package handlers

import (
	"net/http"
	"university_service/handlers/utilities"
)

// TODO: optional fix: Sometimes getting duplicate universities from real api
func handleGetUniInfo(w http.ResponseWriter, r *http.Request) {
	uniApiUrl := utilities.UniversitiesAPIurl + utilities.UniversitiesSearch
	uniApiName := "university"

	name := utilities.GetParamFromRequestURL(r, 5)
	if name == "" {
		http.Error(w, "Expecting format .../{university_name}",
			http.StatusBadRequest)
		return
	}

	params := make(map[string]string)
	params["name"] = name

	var unisFound []utilities.University

	if !utilities.GetResponseAndPopulateData(w, uniApiUrl, uniApiName,
		&params, &unisFound) {
		return
	}

	if !utilities.MarshalAndDisplayData(w, unisFound) {
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
