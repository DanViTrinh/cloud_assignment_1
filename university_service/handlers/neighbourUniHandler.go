package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	util "university_service/handlers/utilities"
)

func NeighborUniHandler(w http.ResponseWriter, r *http.Request) error {

	urlParts, paramErr := util.GetUrlParts(r.URL.Path, 4, 6)
	if paramErr != nil {
		return util.NewClientError(paramErr, http.StatusBadRequest,
			"expecting .../{country_name}/{university_name}")
	}
	searchCountry := urlParts[0]
	uniName := urlParts[1]

	// getting limit if available
	limitAvailable := false
	urlParams := r.URL.Query()
	limitArr, ok := urlParams["limit"]
	var limitInt int

	if len(limitArr) != 1 {
		ok = false
	}

	if ok {
		// to avoid negative integers
		posInt, err := strconv.ParseUint(limitArr[0], 10, 0)
		if err != nil {
			return util.NewClientError(err,
				http.StatusBadRequest, "Only positive integers for limit")
		}
		limitInt = int(posInt)
		limitAvailable = true
	}

	var borderCountries []util.BorderingCountries

	// getting country bordering countries
	countryApiUrlWithCode := util.CountryAPI +
		util.CountryName + "/" + searchCountry
	base, err := url.Parse(countryApiUrlWithCode)
	if err != nil {
		return err
	}

	fullTextParams := url.Values{"fullText": []string{"true"}}
	base.RawQuery = fullTextParams.Encode()

	err = util.FillDataFromApi(base.String(), &borderCountries)

	if err != nil {
		return err
	}

	// getting common name of the bordering countries
	var foundCountries []util.CountryNames

	for i := 0; i < len(borderCountries[0].BorderingCodes); i++ {

		country := borderCountries[0].BorderingCodes[i]
		countryApiUrlWithCode := util.CountryAPI +
			util.CountryCode + "/" + country

		var singleCountryArray []util.CountryNames

		err = util.FillDataFromApi(countryApiUrlWithCode, &singleCountryArray)

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

		uniApiUrl := util.UniAPI +
			util.UniSearch

		base, err := url.Parse(uniApiUrl)
		if err != nil {
			return err
		}

		nameCountryParams := url.Values{}
		nameCountryParams.Add("name", uniName)
		nameCountryParams.Add("country", country.Name.Common)
		base.RawQuery = nameCountryParams.Encode()

		var foundUnis []util.University
		err = util.FillDataFromApi(base.String(), &foundUnis)

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

	return util.DisplayData(w, &finalUnis)
}
