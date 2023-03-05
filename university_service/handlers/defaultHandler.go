package handlers

import (
	"fmt"
	"net/http"
	util "university_service/handlers/utilities"
)

// Handler for root level of url.
//
// Returns:
//
//	ServerError - if failing to give a response
func EmptyHandler(w http.ResponseWriter, r *http.Request) error {
	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "text/html")

	output := "This service does not provide any functionality on root " +
		"path level. Please use paths " +
		"<a href=\"" + util.UniInfoPath + "\">" +
		util.UniInfoPath +
		"</a> or " +
		"<a href=\"" + util.NeighborUnisPath + "\">" +
		util.NeighborUnisPath +
		"</a> or " +
		"<a href=\"" + util.DiagPath + "\">" +
		util.DiagPath +
		"</a>."

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		return util.NewServerError(err, http.StatusInternalServerError,
			util.InternalErrMsg, util.OutputErrMsg)
	}
	return nil
}
