package utilities

import (
	"fmt"
	"net/http"
)

// Adds extra country info to unis like language and map location.
//
// Parameters:
//
//	unis - pointer to the uni array that is gonna get extra country info
//
// Returns:
//
//	ServerError - if failed to fill data from api
func AddCountryInfoToUnis(unis *[]Uni) error {
	//TODO: change implementation, weird.
	//TODO: make a global foundCountries to lessen api calls
	//PROBLEM: the response from the api is a single item array
	foundCountries := make(map[string][]Country)
	for index, uni := range *unis {

		singleCountryArray, ok := foundCountries[uni.IsoCode]

		if ok {
			(*unis)[index].Languages = singleCountryArray[0].Languages
			(*unis)[index].Languages = singleCountryArray[0].Languages
			(*unis)[index].Map = singleCountryArray[0].Maps.Map
		} else {
			countryApiUrlWithCode := CountryAPI +
				CountryCode + "/" + uni.IsoCode

			var singleCountryArray []Country

			err := FillCountriesWithURL(countryApiUrlWithCode, &singleCountryArray)
			if err != nil {
				return err
			}

			(*unis)[index].Languages = singleCountryArray[0].Languages
			(*unis)[index].Map = singleCountryArray[0].Maps.Map

			foundCountries[uni.IsoCode] = singleCountryArray
		}
	}
	return nil

}

// Fills data from uni api. Uni api returns an empty array if nothing is found
// so a separate function is required.
//
// Parameters:
//
//	apiUrl - the url of the api that gets data
//	data - the data that will be filled
//
// Returns:
//
//	ServerError - if the call to api fails or fails during marshal of data
func FillUnisWithURL(apiURL string, unisToFil *[]Uni) error {
	res, err := http.Get(apiURL)
	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, "error in getting response from "+apiURL+" :")
	}

	// if status code is not ok
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("got status: %d from %s", res.StatusCode, apiURL)
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, "error in getting response from "+apiURL+" :")
	}

	return FillDataWithRes(res, unisToFil)
}

// Fills empty country array from api. Country api returns http.StatusNotFound
// when nothing is found so a separate function is required.
//
// Parameters:
//
//	apiUrl - the url of the api that gets data
//	data - the data that will be filled
//
// Returns:
//
//	ServerError - if the call to api fails or fails during marshal of data
func FillCountriesWithURL(apiURL string, countries *[]Country) error {
	res, err := http.Get(apiURL)
	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, "error in getting response from "+apiURL+" :")
	}

	// no country was found
	if res.StatusCode == http.StatusNotFound {
		// set countries to an empty country
		countries = &[]Country{}
		return nil
	}

	// if status code is not ok
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("got status: %d from %s", res.StatusCode, apiURL)
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, "error in getting response from "+apiURL+" :")
	}

	return FillDataWithRes(res, countries)
}
