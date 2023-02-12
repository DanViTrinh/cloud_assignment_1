package handlers

import (
	"net/http"
	"university_service/handlers/utilities"
)

// TODO: optional fix: Sometimes getting duplicate universities from real api
func handleGetUniInfo(w http.ResponseWriter, r *http.Request) {
	uniUrl := UniversitiesAPIurl + UniversitiesSearch
	uniName := "university"

	name := utilities.GetParamFromRequestURL(r, 5)
	if name == "" {
		http.Error(w, "Expecting format .../{university_name}",
			http.StatusBadRequest)
		return
	}

	params := make(map[string]string)
	params["name"] = name

	res := utilities.GetResponseFromApi(w, uniUrl, uniName, params)
	if res == nil {
		return
	}

	var unisFound []University

	if !utilities.PopulateDataWithResponse(w, res, uniName, &unisFound) {
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
