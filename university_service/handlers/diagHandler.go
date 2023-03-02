package handlers

import (
	"net/http"
	"time"
	u "university_service/handlers/utilities"
)

var startTime time.Time = time.Now()

func uptimeInSeconds() time.Duration {
	return time.Since(startTime) / u.NanoSecondsInAsecond
}

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

	diagInfo.Uptime = uptimeInSeconds()

	return u.MarshalAndDisplayData(w, diagInfo)
}
