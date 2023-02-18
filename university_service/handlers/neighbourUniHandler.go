package handlers

import (
	"net/http"
	"university_service/handlers/utilities"
)

func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) error {
	var borderCountries []utilities.BorderingCountries

	countryApiUrlWithCode := utilities.CountriesAPIurl +
		utilities.CountriesAlphaCode + "/TEST"
	err := utilities.GetResponseAndPopulateData(countryApiUrlWithCode,
		nil, &borderCountries)
	if err != nil {
		return err
	}

	return utilities.MarshalAndDisplayData(w, &borderCountries)
}
