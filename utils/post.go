package utils

import (
	"io"
	"bytes"
	"encoding/json"
	"net/http"
)

func POST[T any](url string, body any, contentType string, response *T) (int, error) {

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(url, contentType, bytes.NewReader(bodyBytes))
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	// Parse the JSON response
	err = json.Unmarshal(responseBody, response)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}