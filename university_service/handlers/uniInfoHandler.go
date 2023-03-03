package handlers

import (
	"errors"
	"net/http"
	"net/url"
	util "university_service/handlers/utilities"
)

// TODO: optional fix: Sometimes getting duplicate universities from real api
func handleGetUniInfo(w http.ResponseWriter, r *http.Request) error {

	name, err := util.GetParamFromRequestURL(r, 5)
	if err != nil {
		return util.NewClientError(err, http.StatusBadRequest,
			"expecting format .../{university_name}")
	}

	uniApiUrl, err := url.Parse(util.UniAPI + util.UniSearch)
	if err != nil {
		return err
	}

	params := url.Values{"name": []string{name}}
	uniApiUrl.RawQuery = params.Encode()

	var unisFound []util.University

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
		// return util.NewRestErrorWrapper(fmt.Errorf("%s %s", r.Method,
		// 	util.NotImplementedMsg),
		// 	http.StatusNotImplemented, util.NotImplementedMsg,
		// 	util.UnsensitiveServerError)
		userErrMessage := r.Method + " " + util.NotImplementedMsg
		return util.NewServerError(errors.New(userErrMessage),
			http.StatusInternalServerError, userErrMessage, userErrMessage)
	}
}
