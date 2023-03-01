package handlers

import (
	"fmt"
	"net/http"
	"strings"
	util "university_service/handlers/utilities"
)

func GetCountryAndUniName(r *http.Request) (map[string]string, error) {
	parts := strings.Split(r.URL.Path, "/")
	desiredLen := 6

	params := make(map[string]string)

	//TODO FIX ERROR
	if len(parts) != desiredLen {
		return nil, fmt.Errorf("expecting a value at: %d in %s",
			desiredLen, r.URL.Path)
	}

	params["country"] = parts[desiredLen-2]
	params["name"] = parts[desiredLen-1]
	return params, nil

}

func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) error {
	params, err1 := GetCountryAndUniName(r)
	//TODO
	if err1 != nil {
		return err1
	}
	searchCountry := params["country"]
	uniName := params["name"]

	var borderCountries []util.BorderingCountries

	countryApiUrlWithCode := util.CountriesAPIurl +
		util.CountriesName + "/" + searchCountry
	err := util.GetResponseAndPopulateData(countryApiUrlWithCode,
		nil, &borderCountries)
	if err != nil {
		return util.NewRestErrorWrapper(err, http.StatusInternalServerError,
			"Could not get neighbours from countries api", util.ServerError)
	}

	var foundCountries []util.CountryName

	for _, country := range borderCountries[0].BorderingCodes {

		countryApiUrlWithCode := util.CountriesAPIurl +
			util.CountriesAlphaCode + "/" + country

		var singleCountryArray []util.CountryName

		err = util.GetResponseAndPopulateData(countryApiUrlWithCode, nil,
			&singleCountryArray)

		// TODO HANDLE PROPER ERRROR
		if err != nil {
			return err
		}

		foundCountries = append(foundCountries, singleCountryArray[0])

	}

	var finalUnis []util.University
	for _, country := range foundCountries {

		uniApiUrl := util.UniversitiesAPIurl +
			util.UniversitiesSearch
		params := make(map[string]string)

		params["name"] = uniName
		params["country"] = country.Name.Common
		var foundUnis []util.University
		err := util.GetResponseAndPopulateData(uniApiUrl,
			&params, &foundUnis)

		// TODO
		if err != nil {
			return err
		}

		finalUnis = append(finalUnis, foundUnis...)
	}

	return util.MarshalAndDisplayData(w, &finalUnis)
}
