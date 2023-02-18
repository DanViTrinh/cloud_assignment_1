package handlers

import (
	"errors"
	"log"
	"net/http"
	util "university_service/handlers/utilities"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

// TODO: add better error handling
func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}

	var restErr util.RestErrorWraper

	if errors.As(err, &restErr) {
		switch restErr.ErrorKind {
		case util.ClientError:
			http.Error(w, restErr.Message, restErr.StatusCode)
		case util.UnsensitiveServerError:
			http.Error(w, restErr.Message, restErr.StatusCode)
		default:
			log.Println(restErr.Message)
			http.Error(w, util.StandardInternalServerErrorMsg,
				restErr.StatusCode)
		}
		log.Printf("Cause of Error: %v", restErr.OriginalError)
		return
	}

	log.Printf("Unexpected error type, original error: %v", err)
	http.Error(w, util.StandardInternalServerErrorMsg,
		http.StatusInternalServerError)
}
