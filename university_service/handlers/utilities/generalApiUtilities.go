package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func populateDataWithResponse(res *http.Response, data interface{}) error {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return NewRestErrorWrapper(err, http.StatusInternalServerError,
			"error during reading response", ServerError)
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return NewRestErrorWrapper(err, http.StatusInternalServerError,
			"error during unmarshaling ", ServerError)
		// http.Error(w, "Error during unmarshaling from ",
		// 	http.StatusInternalServerError)
	}
	return nil
}

// TODO: consider using get instead
func getResponseFromApi(apiURL string,
	params *map[string]string) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, NewRestErrorWrapper(err, http.StatusInternalServerError,
			"error in creating new request to "+apiURL, ServerError)
		// http.Error(w, "Error in creating new request to "+apiURL+" API",
		// 	http.StatusInternalServerError)
	}

	if params != nil {
		q := request.URL.Query()

		for key, val := range *params {
			q.Add(key, val)
		}

		request.URL.RawQuery = q.Encode()
	}

	// request.Header.Add("content-type", "application/json")
	client := &http.Client{}
	defer client.CloseIdleConnections()

	res, err := client.Do(request)

	if err != nil {
		return nil, NewRestErrorWrapper(err, http.StatusInternalServerError,
			"error in getting response from "+apiURL, ServerError)
		// http.Error(w, "Error in getting response from "+apiName+" API",
		// 	http.StatusInternalServerError)
		// return nil
	}
	return res, nil
}

func GetResponseAndPopulateData(apiURL string,
	params *map[string]string, data interface{}) error {

	res, err := getResponseFromApi(apiURL, params)
	if err != nil {
		return err
	}

	return populateDataWithResponse(res, &data)
}

func MarshalAndDisplayData(w http.ResponseWriter, data interface{}) error {

	w.Header().Add("content-type", "application/json")

	jsonEncodedData, err := json.Marshal(data)

	if err != nil {
		return NewRestErrorWrapper(err, http.StatusInternalServerError,
			"error during marshalling data", ServerError)
		// http.Error(w, "Error during marshalling data",
		// 	http.StatusInternalServerError)
	}

	_, err = fmt.Fprint(w, string(jsonEncodedData))
	if err != nil {
		return NewRestErrorWrapper(err, http.StatusInternalServerError,
			"error during writing response", ServerError)
		// http.Error(w, "Error during writing response",
		// 	http.StatusInternalServerError)
	}
	return nil
}

func GetParamFromRequestURL(r *http.Request, desiredLen int) (string, error) {
	parts := strings.Split(r.URL.Path, "/")

	if (len(parts) == desiredLen && parts[desiredLen-1] != "") ||
		(len(parts) == desiredLen+1 && parts[desiredLen-1] != "" &&
			parts[desiredLen] == "") {
		return parts[desiredLen-1], nil
	}
	return "", fmt.Errorf("expecting a value at: %d in %s",
		desiredLen, r.URL.Path)
	// return "", errors.New("Expecting a value at: " + desiredLen + "in " )
	// string(r.URL.Path))
}
