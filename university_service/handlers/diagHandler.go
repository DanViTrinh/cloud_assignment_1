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

	respUni, err := u.GetResponseFromApi(u.UniversitiesAPIurl, nil)
	if err != nil {
		return err
	}
	diagInfo.UniApiStatus = respUni.Status

	respCountry, err := u.GetResponseFromApi(u.UniversitiesAPIurl, nil)
	if err != nil {
		return err
	}
	diagInfo.CountryApiStatus = respCountry.Status

	diagInfo.Version = u.Version

	diagInfo.Uptime = time.Since(startTime) / u.NanoSecondsInAsecond

	return u.MarshalAndDisplayData(w, diagInfo)
}
