package handlers

import (
	"net/http"
	"time"
	u "university_service/handlers/utilities"
)

// Starts time at the beggining of service
var startTime time.Time = time.Now()

// TODO: Debate: what to do when the service is down and unavailable
// should i set a default status code instead of returning if error occured
func DiagHandler(w http.ResponseWriter, r *http.Request) error {

	var diagInfo u.DiagInfo

	respUni, uniErr := http.Get(u.UniAPI)
	if uniErr != nil {
		return u.NewRestErrorWrapper(uniErr, http.StatusInternalServerError,
			"error in getting response from "+u.UniAPI, u.ServerError)
	}
	diagInfo.UniApiStatus = respUni.Status

	// TODO: CHANGE ERROR
	respCountry, err := http.Get(u.CountryAPI)
	if err != nil {
		return err
	}
	diagInfo.CountryApiStatus = respCountry.Status

	diagInfo.Version = u.Version

	diagInfo.Uptime = time.Since(startTime) / u.NanoSecondsInAsecond

	return u.DisplayData(w, diagInfo)
}
