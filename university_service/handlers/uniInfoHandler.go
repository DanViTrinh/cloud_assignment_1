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

	err = util.AddCountryInfoToUnis(&unisFound)
	//TODO
	if err != nil {
		return err
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
