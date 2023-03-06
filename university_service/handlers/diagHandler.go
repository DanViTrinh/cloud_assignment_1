package handlers

import (
	"net/http"
	"time"
	util "university_service/utilities"
)

// Starts time at the beginning of service
var startTime time.Time = time.Now()

// Diagnostic handler. Gets status of foreign api's and current api
//
// Returns:
//
//	ServerError - if contact with foreign api or writing response failed
func DiagHandler(w http.ResponseWriter, r *http.Request) error {

	var diagInfo util.DiagInfo

	respUni, uniErr := http.Get(util.UniAPI)
	if uniErr != nil {
		return util.NewServerError(uniErr, http.StatusInternalServerError,
			util.InternalErrMsg, "error in getting response from uni api")
	}
	diagInfo.UniApiStatus = respUni.Status

	TestCountryUrl := util.CountryAPI + util.CountryCode + util.TestCountryCode
	respCountry, countryErr := http.Get(TestCountryUrl)
	if countryErr != nil {
		return util.NewServerError(countryErr, http.StatusInternalServerError,
			util.InternalErrMsg, "error in getting response from country api")
	}
	diagInfo.CountryApiStatus = respCountry.Status

	diagInfo.Version = util.Version

	diagInfo.Uptime = time.Since(startTime) / util.NanoSecInSec

	return util.DisplayData(w, diagInfo)
}
