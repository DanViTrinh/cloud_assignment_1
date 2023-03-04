package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Gets response from api url Populates data with the response
func FillDataFromApi(apiURL string, data interface{}) error {
	res, err := http.Get(apiURL)
	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, "error in getting response from api")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return NewServerError(err, http.StatusInternalServerError,
			InternalErrMsg, "error during reading response")
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
// startParam is the index of the param that you want to start with.
// amount is the amount of params you want out.
// amount has to be amount from startParam to end of param
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
