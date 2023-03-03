package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Gets response from api url Populates data with the response
func FillDataFromApi(apiURL string, data interface{}) error {
	res, err := http.Get(apiURL)
	if err != nil {
		return err
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

// TODO: necessary?
func GetParamFromRequestURL(r *http.Request, desiredLen int) (string, error) {
	parts := strings.Split(r.URL.Path, "/")

	if (len(parts) == desiredLen && parts[desiredLen-1] != "") ||
		(len(parts) == desiredLen+1 && parts[desiredLen-1] != "" &&
			parts[desiredLen] == "") {
		return parts[desiredLen-1], nil
	}
	return "", fmt.Errorf("expecting a value at: %d in %s",
		desiredLen, r.URL.Path)
}
