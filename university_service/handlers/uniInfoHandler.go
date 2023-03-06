package handlers

import (
	"errors"
	"net/http"
	"net/url"
	util "university_service/utilities"
)

// Handles getting uni info
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

// Handles get request for uni info.
// Displays gives found unis as a response
func handleGetUniInfo(w http.ResponseWriter, r *http.Request) error {

	// Get the url params
	urlParams, err := util.GetUrlParts(r.URL.Path, 4, 1)
	if err != nil {
		return util.NewClientError(err, http.StatusBadRequest,
			"expecting format .../{university_name}")
	}

	// parsing uni url to get a URL
	uniApiUrl, err := url.Parse(util.UniAPI + util.UniSearch)
	if err != nil {
		return err
	}

	// Adding params to url
	params := url.Values{util.UniNameParam: []string{urlParams[0]}}
	uniApiUrl.RawQuery = params.Encode()

	var unisFound []util.Uni

	// filling unisFound
	err = util.FillUnisWithURL(uniApiUrl.String(), &unisFound)
	if err != nil {
		return err
	}

	// add extra info to unis
	err = util.AddCountryInfoToUnis(&unisFound)

	if err != nil {
		return err
	}

	// gives response
	return util.DisplayData(w, unisFound)
}
