package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FillDataWithRes(res *http.Response, data interface{}) error {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, "error when reading response")
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, UnmarshalErrMsg)
	}
	return nil

}

// Display a struct through writer
// The displayed struct will be displayed as json
//
// Parameters:
//
//	w - the http.ResponseWriter that the data will be displayed to
//	data - the struct that is gonna be displayed
//
// Returns:
//
//	ServerError - if fails during marshal of data or when printing response
func DisplayData(w http.ResponseWriter, data interface{}) error {

	w.Header().Add("content-type", "application/json")

	jsonEncodedData, err := json.Marshal(data)

	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, UnmarshalErrMsg)
	}

	_, err = fmt.Fprint(w, string(jsonEncodedData))
	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, ResponseErrMsg)
	}
	return nil
}

// Gets params from url from startParam to the end of param
// has to check for a matching length. Error will be returned if empty param
// is found.
//
// Parameters:
//
//	url - the url to fetch parts from. Should not contain parameters ie. "?"
//	startParam - the index of the param that you want to start with.
//	amount - amount of params you want returned. amount has to be amount from
//			 startParam to end of param
//
// Returns:
//
//	string[] - list of url parts from startParam to end
//	error - if the url has invalid length or empty in desired parts
func GetUrlParts(url string, startParam, amount int) ([]string, error) {

	// removes trailing "/"
	path := strings.TrimSuffix(url, "/")

	// checks for empty params
	// if it contains "//" there must be one parameter that is empty
	if !strings.Contains(path, "//") {
		parts := strings.Split(strings.TrimSuffix(url, "/"), "/")
		// checks valid length
		if len(parts) == startParam+amount {
			// returns desired params
			return parts[startParam:], nil
		}
	}

	return nil, errors.New("invalid url length or empty url params")
}
