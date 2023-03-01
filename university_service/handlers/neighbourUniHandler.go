package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	util "university_service/handlers/utilities"
)

func GetCountryUniAndLimit(r *http.Request) (map[string]string, error) {
	parts := strings.Split(r.URL.Path, "/")
	desiredLen := 6

	params := make(map[string]string)

	//TODO FIX ERROR
	if len(parts) != desiredLen {
		return nil, fmt.Errorf("expecting a value at: %d in %s",
			desiredLen, r.URL.Path)
	}

	params["country"] = parts[desiredLen-2]

	if r.URL.Query().Has("limit") {
		params["name"] = strings.Split(parts[desiredLen-1], "?")[0]
		params["limit"] = r.URL.Query().Get("limit")
	} else {
		params["name"] = parts[desiredLen-1]
	}
	return params, nil

}

func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) error {
	params, err1 := GetCountryUniAndLimit(r)
	//TODO
	if err1 != nil {
		return err1
	}
	searchCountry := params["country"]
	uniName := params["name"]
	limit, ok := params["limit"]

	// gettting the limit if available
	var limitInt int
	limitAvailable := false
	if ok {
		// to avoid negative integers
		//TODO use atoi and check for neg vals
		posInt, err := strconv.ParseUint(limit, 10, 0)
		if err != nil {
			return util.NewRestErrorWrapper(err,
				http.StatusBadRequest, "Only positive integers for limit ",
				util.ClientError)
		}
		// todo try avoiding casting
		limitInt = int(posInt)
		limitAvailable = true
	}

	var borderCountries []util.BorderingCountries

	// getting country bordering countries
	countryApiUrlWithCode := util.CountriesAPIurl +
		util.CountriesName + "/" + searchCountry
	fullTextParam := map[string]string{
		"fullText": "true",
	}

	err := util.GetResponseAndPopulateData(countryApiUrlWithCode,
		&fullTextParam, &borderCountries)

	if err != nil {
		return util.NewRestErrorWrapper(err, http.StatusInternalServerError,
			"Could not get neighbours from countries api", util.ServerError)
	}

	// getting commmon name of the bordering countries
	var foundCountries []util.CountryName

	for i := 0; i < len(borderCountries[0].BorderingCodes); i++ {

		country := borderCountries[0].BorderingCodes[i]
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

	// getting universities in bordering countries
	var finalUnis []util.University
	// for _, country := range foundCountries {

	for i := 0; i < len(foundCountries) &&
		(len(finalUnis) <= limitInt || !limitAvailable); i++ {

		country := foundCountries[i]

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

	// limiting the final array
	if limitAvailable && len(finalUnis) > limitInt {
		finalUnis = finalUnis[:limitInt]
	}

	// adding country information to  uni
	// TODO debate adding the already found countries to the countries cache
	err = util.AddCountryInfoToUnis(&finalUnis)
	//TODO
	if err != nil {
		return err
	}

	// TODO MAKE PROPER LIMIT TO REDUCE API CALLS LATER

	return util.MarshalAndDisplayData(w, &finalUnis)
}
