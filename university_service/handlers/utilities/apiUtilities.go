package utilities

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
	foundCountries := make(map[string][]MissingFieldsFromCountry)
	for index, uni := range *unis {

		singleCountryArray, ok := foundCountries[uni.IsoCode]

		if ok {
			(*unis)[index].Languages = singleCountryArray[0].Languages
			(*unis)[index].Languages = singleCountryArray[0].Languages
			(*unis)[index].Map =
				singleCountryArray[0].Maps[DesiredMap]
		} else {
			countryApiUrlWithCode := CountryAPI +
				CountryCode + "/" + uni.IsoCode

			var singleUniArray []MissingFieldsFromCountry

			err := FillDataFromApi(countryApiUrlWithCode, &singleUniArray)
			if err != nil {
				return err
			}

			(*unis)[index].Languages = singleUniArray[0].Languages
			(*unis)[index].Map = singleUniArray[0].Maps[DesiredMap]

			foundCountries[uni.IsoCode] = singleUniArray
		}
	}
	return nil

}
