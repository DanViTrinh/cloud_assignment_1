package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	util "university_service/handlers/utilities"
)

// Handles call neighbor unis
func NeighborUniHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return handleGetNeighborUni(w, r)
	default:
		userErrMessage := r.Method + " " + util.NotImplementedMsg
		return util.NewServerError(errors.New(userErrMessage),
			http.StatusInternalServerError, userErrMessage, userErrMessage)
	}
}

// Handles get request to neighbor unis.
func handleGetNeighborUni(w http.ResponseWriter, r *http.Request) error {
	// get param parts
	urlParts, paramErr := util.GetUrlParts(r.URL.Path, 4, 2)
	if paramErr != nil {
		return util.NewClientError(paramErr, http.StatusBadRequest,
			"expecting .../{country_name}/{university_name}")
	}
	searchCountry := urlParts[0]
	uniName := urlParts[1]

	// get limit if available
	limit, limitAvailable, limitErr := getLimit(r)
	if limitErr != nil {
		return limitErr
	}

	// get border codes
	borderCountries, borderErr := getBorderCodes(searchCountry)
	if borderErr != nil {
		return borderErr
	}

	// get country names for borders
	foundCountries, foundCountryErr := getCountryNames(borderCountries)
	if foundCountryErr != nil {
		return foundCountryErr
	}

	// getting universities in bordering countries
	var finalUnis []util.Uni
	uniApiUrl, err := url.Parse(util.UniAPI +
		util.UniSearch)
	if err != nil {
		return err
	}
	for i := 0; i < len(foundCountries) &&
		(len(finalUnis) <= limit || !limitAvailable); i++ {

		country := foundCountries[i]
		var foundUnis []util.Uni
		err := fillUnisFromCountry(country, uniName, &foundUnis, *uniApiUrl)
		if err != nil {
			return err
		}
		finalUnis = append(finalUnis, foundUnis...)
	}

	// limiting the final array
	if limitAvailable && len(finalUnis) > limit {
		finalUnis = finalUnis[:limit]
	}

	// adding country information to  uni
	// TODO debate adding the already found countries to the countries cache
	addErr := util.AddCountryInfoToUnis(&finalUnis)
	if addErr != nil {
		return addErr
	}

	return util.DisplayData(w, &finalUnis)
}

// Gets the limit param from url
//
// example use case:
//
//	limit , ok , err := getLimit(request)
//
// limit param must be a positive int
// if limit param is invalid error will return
// if limit not available false will be returned (ok will be false)
//
// Returns:
//
// ClientError - is returned when the url param is invalid
func getLimit(r *http.Request) (int, bool, error) {
	limitAvailable := false
	urlParams := r.URL.Query()
	limitArr, ok := urlParams["limit"]
	var limit int

	if len(limitArr) != 1 {
		ok = false
	}

	if ok {
		// to avoid negative integers
		posInt, err := strconv.ParseUint(limitArr[0], 10, 0)
		if err != nil {
			return limit, false, util.NewClientError(err,
				http.StatusBadRequest, "Only positive integers for limit")
		}
		limit = int(posInt)
		limitAvailable = true
	}
	return limit, limitAvailable, nil
}

// Returns border codes for searchCountry
//
// Params:
//
//	searchCountry - the country to search for borders
//
// Returns:
//
//	[]string - a list of the bordering country codes
//	ServerError - if the operation failed
func getBorderCodes(searchCountry string) ([]string, error) {
	var borderCountries []util.Country

	countryApiUrlWithCode := util.CountryAPI +
		util.CountryName + "/" + searchCountry
	base, err := url.Parse(countryApiUrlWithCode)
	if err != nil {
		return nil, err
	}

	fullTextParams := url.Values{"fullText": []string{"true"}}
	base.RawQuery = fullTextParams.Encode()

	err = util.FillCountriesWithURL(base.String(), &borderCountries)

	if err != nil {
		return nil, err
	}
	return borderCountries[0].BorderingCodes, nil
}

// Getting common name of the bordering countries
//
// Parameters:
//
//	countryCodes - the country codes used to find country names
//
// Returns:
//
//	[]CountryNames - list of country names for each country code
//					 country names will have the same index as country code
//	ServerError - if the operation failed
func getCountryNames(countryCodes []string) ([]util.Country, error) {
	var foundCountries []util.Country

	for i := 0; i < len(countryCodes); i++ {

		country := countryCodes[0]
		countryApiUrlWithCode := util.CountryAPI +
			util.CountryCode + "/" + country

		var singleCountryArray []util.Country

		err := util.FillCountriesWithURL(countryApiUrlWithCode, &singleCountryArray)

		if err != nil {
			return nil, err
		}

		foundCountries = append(foundCountries, singleCountryArray[0])
	}
	return foundCountries, nil
}

// Fill unis array with unis from a country
//
// Parameters:
//
//	country - the country Name/names to find unis from
//	uniName - the uni name for searching up uni
//	unis - a pointer to the unis array to be filled up
func fillUnisFromCountry(country util.Country, uniName string,
	unis *[]util.Uni, apiUrl url.URL) error {

	nameCountryParams := url.Values{}
	nameCountryParams.Add("name", uniName)
	nameCountryParams.Add("country", country.Name.Common)
	apiUrl.RawQuery = nameCountryParams.Encode()

	err := util.FillUnisWithURL(apiUrl.String(), unis)
	if err != nil {
		return err
	}
	return nil

}
