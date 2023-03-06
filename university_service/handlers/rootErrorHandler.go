package handlers

import (
	"log"
	"net/http"
	util "university_service/utilities"
)

// Wraps all other handlers with this func to be able to ServeHttp
type RootHandler func(http.ResponseWriter, *http.Request) error

// Required function to use http.Handle
// By using this errors can be handled in one place.
// Handles all error here.
func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// calls the original function, and all errors will return to this level
	err := fn(w, r)
	if err == nil {
		return
	}

	// handle error:

	// checks for error type
	switch e := err.(type) {
	case util.ClientError:
		log.Println("Client error:")
		http.Error(w, e.Message, e.StatusCode)
	case util.ServerError:
		log.Println(e.DevMessage)
		http.Error(w, e.UsrMessage, e.StatusCode)
	default:
		log.Println("Non rest error:")
		http.Error(w, util.InternalErrMsg, http.StatusInternalServerError)
	}
	// logs the original error
	log.Println("\t" + err.Error())
}
