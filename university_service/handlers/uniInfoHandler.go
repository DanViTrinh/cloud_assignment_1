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

	//TODO: make a general function for neighbour uni
	//TODO: change implementation, weird.
	//PROBLEM: the response from the api is a single item array
	countryApiName := "rest countries"
	foundCountries := make(map[string][]utilities.MissingFieldsFromCountry)
	for index, uni := range unisFound {

		singleCountryArray, ok := foundCountries[uni.IsoCode]

		if ok {
			unisFound[index].Languages = singleCountryArray[0].Languages
			unisFound[index].Map =
				singleCountryArray[0].Maps[utilities.DesiredMap]
		} else {
			countryApiUrlWithCode := utilities.CountriesAPIurl +
				utilities.CountriesAlphaCode + "/" + uni.IsoCode

			var singleUniArray []utilities.MissingFieldsFromCountry

			if !utilities.GetResponseAndPopulateData(w, countryApiUrlWithCode,
				countryApiName, nil, &singleUniArray) {
				return
			}
			unisFound[index].Languages = singleUniArray[0].Languages
			unisFound[index].Map = singleUniArray[0].Maps[utilities.DesiredMap]

			foundCountries[uni.IsoCode] = singleUniArray
		}
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
