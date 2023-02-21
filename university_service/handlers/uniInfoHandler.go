package handlers

import (
	"fmt"
	"net/http"
	util "university_service/handlers/utilities"
)

// TODO: optional fix: Sometimes getting duplicate universities from real api
func handleGetUniInfo(w http.ResponseWriter, r *http.Request) error {
	uniApiUrl := util.UniversitiesAPIurl + util.UniversitiesSearch

	name, err := util.GetParamFromRequestURL(r, 5)
	if err != nil {
		return util.NewRestErrorWrapper(err, http.StatusBadRequest,
			"expecting format .../{university_name}", util.ClientError)
		// http.Error(w, "Expecting format .../{university_name}",
		// 	http.StatusBadRequest)
		// return
	}

	params := make(map[string]string)
	params["name"] = name

	var unisFound []util.University

	err = util.GetResponseAndPopulateData(uniApiUrl, &params, &unisFound)
	if err != nil {
		return err
	}

	//TODO: make a general function for neighbour uni
	//TODO: change implementation, weird.
	//TODO: make a global foundCountries to lessen api calls
	//PROBLEM: the response from the api is a single item array
	foundCountries := make(map[string][]util.MissingFieldsFromCountry)
	for index, uni := range unisFound {

		singleCountryArray, ok := foundCountries[uni.IsoCode]

		if ok {
			unisFound[index].Languages = singleCountryArray[0].Languages
			unisFound[index].Map =
				singleCountryArray[0].Maps[util.DesiredMap]
		} else {
			countryApiUrlWithCode := util.CountriesAPIurl +
				util.CountriesAlphaCode + "/" + uni.IsoCode

			var singleUniArray []util.MissingFieldsFromCountry

			err = util.GetResponseAndPopulateData(countryApiUrlWithCode, nil,
				&singleUniArray)
			if err != nil {
				return err
			}
			// if !util.GetResponseAndPopulateData(w, countryApiUrlWithCode,
			// 	countryApiName, nil, &singleUniArray) {
			// 	return
			// }
			unisFound[index].Languages = singleUniArray[0].Languages
			unisFound[index].Map = singleUniArray[0].Maps[util.DesiredMap]

			foundCountries[uni.IsoCode] = singleUniArray
		}
	}

	return util.MarshalAndDisplayData(w, unisFound)
}

func UniInfoHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return handleGetUniInfo(w, r)
	default:
		return util.NewRestErrorWrapper(fmt.Errorf("%s %s", r.Method,
			util.NotImplementedMsg),
			http.StatusNotImplemented, util.NotImplementedMsg,
			util.UnsensitiveServerError)
	}
}
