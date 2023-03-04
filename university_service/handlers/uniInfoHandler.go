package handlers

import (
	"errors"
	"net/http"
	"net/url"
	util "university_service/handlers/utilities"
)

// TODO: optional fix: Sometimes getting duplicate universities from real api
func handleGetUniInfo(w http.ResponseWriter, r *http.Request) error {

	urlParams, err := util.GetUrlParts(r.URL.Path, 4, 1)
	if err != nil {
		return util.NewClientError(err, http.StatusBadRequest,
			"expecting format .../{university_name}")
	}

	uniApiUrl, err := url.Parse(util.UniAPI + util.UniSearch)
	if err != nil {
		return err
	}

	params := url.Values{"name": []string{urlParams[0]}}
	uniApiUrl.RawQuery = params.Encode()

	var unisFound []util.Uni

	err = util.FillDataFromApi(uniApiUrl.String(), &unisFound)
	if err != nil {
		return err
	}

	err = util.AddCountryInfoToUnis(&unisFound)
	//TODO
	if err != nil {
		return err
	}

	return util.DisplayData(w, unisFound)
}

func UniInfoHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return handleGetUniInfo(w, r)
	default:
		userErrMessage := r.Method + " " + util.NotImplementedMsg
		return util.NewServerError(errors.New(userErrMessage),
			http.StatusInternalServerError, userErrMessage, userErrMessage)
	}
}
