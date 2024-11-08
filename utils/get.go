package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func GET[T any](url string, response *T) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	// Parse the JSON response
	err = json.Unmarshal(body, response)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}