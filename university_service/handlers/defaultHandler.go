package handlers

import (
	"fmt"
	"net/http"
	"university_service/handlers/utilities"
)

func EmptyHandler(w http.ResponseWriter, r *http.Request) error {
	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "text/html")

	output := "This service does not provide any functionality on root " +
		"path level. Please use paths " +
		"<a href=\"" + utilities.UniInfoPath + "\">" +
		utilities.UniInfoPath +
		"</a> or " +
		"<a href=\"" + utilities.NeighbourUnisPath + "\">" +
		utilities.NeighbourUnisPath +
		"</a> or " +
		"<a href=\"" + utilities.DiagPath + "\">" +
		utilities.DiagPath +
		"</a>."

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error when returning output",
			http.StatusInternalServerError)
	}

}
