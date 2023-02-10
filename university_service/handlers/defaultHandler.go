package handlers

import (
	"fmt"
	"net/http"
)

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "text/html")

	output := "This service does not provide any functionality on root path level. Please use paths " +
		"<a href=\"" + UniInfoPath + "\">" + UniInfoPath + "</a> or " +
		"<a href=\"" + NeighbourUnisPath + "\">" + NeighbourUnisPath + "</a> or " +
		"<a href=\"" + DiagPath + "\">" + DiagPath + "</a>."

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}

}
