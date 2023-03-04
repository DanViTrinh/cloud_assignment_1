package handlers

import (
	"net/http"
	"time"
	u "university_service/handlers/utilities"
)

// Starts time at the beginning of service
var startTime time.Time = time.Now()

// TODO: Debate: what to do when the service is down and unavailable
// should i set a default status code instead of returning if error occurred
func DiagHandler(w http.ResponseWriter, r *http.Request) error {

	var diagInfo u.DiagInfo

	respUni, uniErr := http.Get(u.UniAPI)
	if uniErr != nil {
		return u.NewServerError(uniErr, http.StatusInternalServerError,
			u.InternalErrMsg, "error in getting response from uni api")
	}
	diagInfo.UniApiStatus = respUni.Status

	TestCountryUrl := u.CountryAPI + u.CountryCode + u.TestCountryCode
	respCountry, countryErr := http.Get(TestCountryUrl)
	if countryErr != nil {
		return u.NewServerError(countryErr, http.StatusInternalServerError,
			u.InternalErrMsg, "error in getting response from country api")
	}
	diagInfo.CountryApiStatus = respCountry.Status

	diagInfo.Version = u.Version

	diagInfo.Uptime = time.Since(startTime) / u.NanoSecInSec

	return u.DisplayData(w, diagInfo)
}
