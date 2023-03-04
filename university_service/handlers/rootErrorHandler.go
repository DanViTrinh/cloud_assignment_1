package handlers

import (
	"log"
	"net/http"
	util "university_service/handlers/utilities"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}

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
