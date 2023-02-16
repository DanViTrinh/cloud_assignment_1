package handlers

import (
	"net/http"
	"university_service/handlers/utilities"
)

func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) error {

	var borderCountries []utilities.BorderingCountries

	countryApiUrlWithCode := utilities.CountriesAPIurl +
		utilities.CountriesAlphaCode + "/TEST"
	if !utilities.GetResponseAndPopulateData(w, countryApiUrlWithCode,
		"Country", nil, &borderCountries) {
		return
	}

	if !utilities.MarshalAndDisplayData(w, &borderCountries) {
		return
	}

}
