package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"university_service/utilities"
)

// Handles call neighbor unis
func NeighborUniHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return handleGetNeighborUni(w, r)
	default:
		userErrMessage := r.Method + " " + utilities.NotImplementedMsg
		return utilities.NewServerError(errors.New(userErrMessage),
			http.StatusInternalServerError, userErrMessage, userErrMessage)
	}
}

// Handles get request to neighbor unis.
func handleGetNeighborUni(w http.ResponseWriter, r *http.Request) error {
	// get param parts
	urlParts, paramErr := utilities.GetUrlParts(r.URL.Path, 4, 2)
	if paramErr != nil {
		return utilities.NewClientError(paramErr, http.StatusBadRequest,
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
	borderCodes, borderErr := getBorderCodes(searchCountry)
	if borderErr != nil {
		return borderErr
	}

	// get cca2 of border codes
	borderCountries, countryErr := getCountriesWithCodes(borderCodes)
	if countryErr != nil {
		return countryErr
	}

	// getting universities with uniName
	uniApiUrl, err := url.Parse(utilities.UniAPI +
		utilities.UniSearch)
	if err != nil {
		return err
	}

	uniParams := url.Values{"name": []string{uniName}}
	uniApiUrl.RawQuery = uniParams.Encode()

	var unisWithName []utilities.Uni
	err = utilities.FillUnisWithURL(uniApiUrl.String(), &unisWithName)
	if err != nil {
		return err
	}

	// initializes to empty array
	// if no unis is found empty arr will be displayed
	finalUnis := []utilities.Uni{}

	// loop runs through all the unis with name or until
	// limit amount of unis is found, if limit is available
	for i := 0; i < len(unisWithName) &&
		(len(finalUnis) <= limit || !limitAvailable); i++ {
		found := false
		for bI := 0; bI < len(borderCountries) && !found; bI++ {
			if unisWithName[i].IsoCode == borderCountries[bI].Cca2 {
				finalUnis = append(finalUnis, unisWithName[i])
				found = true
			}
		}

	}

	// limiting the final array
	if limitAvailable && len(finalUnis) > limit {
		finalUnis = finalUnis[:limit]
	}

	// adding country information to  uni
	// TODO debate adding the already found countries to the countries cache
	addErr := utilities.AddCountryInfoToUnis(&finalUnis)
	if addErr != nil {
		return addErr
	}

	return utilities.DisplayData(w, &finalUnis)
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
			return limit, false, utilities.NewClientError(err,
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
	var countryInSingleArr []utilities.Country

	countryApiUrlWithCode := utilities.CountryAPI +
		utilities.CountryName + "/" + searchCountry
	base, err := url.Parse(countryApiUrlWithCode)
	if err != nil {
		return nil, err
	}

	fullTextParams := url.Values{"fullText": []string{"true"}}
	base.RawQuery = fullTextParams.Encode()

	err = utilities.FillCountriesWithURL(base.String(), &countryInSingleArr)

	if err != nil {
		return nil, err
	}
	if len(countryInSingleArr) != 0 {
		return countryInSingleArr[0].BorderingCodes, nil
	}
	// return empty list
	return []string{}, nil
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
func getCountriesWithCodes(countryCodes []string) ([]utilities.Country, error) {
	var foundCountries []utilities.Country

	for i := 0; i < len(countryCodes); i++ {

		country := countryCodes[i]
		countryApiUrlWithCode := utilities.CountryAPI +
			utilities.CountryCode + "/" + country

		var singleCountryArray []utilities.Country

		err := utilities.FillCountriesWithURL(countryApiUrlWithCode, &singleCountryArray)

		if err != nil {
			return nil, err
		}

		foundCountries = append(foundCountries, singleCountryArray[0])
	}
	return foundCountries, nil
}
