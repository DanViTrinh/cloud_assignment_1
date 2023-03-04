package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	util "university_service/handlers/utilities"
)

func NeighborUniHandler(w http.ResponseWriter, r *http.Request) error {

	// get param parts
	urlParts, paramErr := util.GetUrlParts(r.URL.Path, 4, 6)
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

// gets the limit param from url
// example use case:
// limit , ok , err := getLimit(request)
// limit param must be a positive int
// if limit param is invalid error will return
// if limit not available false will be returned (ok will be false)
// client error is returned
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

func getBorderCodes(searchCountry string) ([]string, error) {
	var borderCountries []util.BorderCountries
	// getting country bordering countries
	countryApiUrlWithCode := util.CountryAPI +
		util.CountryName + "/" + searchCountry
	base, err := url.Parse(countryApiUrlWithCode)
	if err != nil {
		return nil, err
	}

	fullTextParams := url.Values{"fullText": []string{"true"}}
	base.RawQuery = fullTextParams.Encode()

	err = util.FillDataFromApi(base.String(), &borderCountries)

	if err != nil {
		return nil, err
	}
	return borderCountries[0].BorderingCodes, nil
}

// getting common name of the bordering countries
func getCountryNames(countryCodes []string) ([]util.CountryNames, error) {
	var foundCountries []util.CountryNames

	for i := 0; i < len(countryCodes); i++ {

		country := countryCodes[0]
		countryApiUrlWithCode := util.CountryAPI +
			util.CountryCode + "/" + country

		var singleCountryArray []util.CountryNames

		err := util.FillDataFromApi(countryApiUrlWithCode, &singleCountryArray)

		if err != nil {
			return nil, err
		}

		foundCountries = append(foundCountries, singleCountryArray[0])
	}
	return foundCountries, nil
}

// fill uni in param with found country unis
func fillUnisFromCountry(country util.CountryNames, uniName string,
	unis *[]util.Uni, apiUrl url.URL) error {

	nameCountryParams := url.Values{}
	nameCountryParams.Add("name", uniName)
	nameCountryParams.Add("country", country.Name.Common)
	apiUrl.RawQuery = nameCountryParams.Encode()

	err := util.FillDataFromApi(apiUrl.String(), &unis)
	if err != nil {
		return err
	}
	return nil

}
