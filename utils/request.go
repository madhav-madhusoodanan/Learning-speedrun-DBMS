package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// HeaderOption is a function type for setting headers
type HeaderOption func(*http.Request)

// WithHeader returns a HeaderOption that sets a single header
func WithHeader(key, value string) HeaderOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// WithHeaders returns a HeaderOption that sets multiple headers
func WithHeaders(headers map[string]string) HeaderOption {
	return func(req *http.Request) {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
}

// WithBearerToken returns a HeaderOption that sets the Authorization header with a bearer token
func WithBearerToken(token string) HeaderOption {
	return func(req *http.Request) {
		req.Header.Set("Authorization", "Bearer "+token)
	}
}

// GET makes a GET request with custom headers and unmarshals the response into the provided type
func REQUEST[T any](url string, response *T, opts ...HeaderOption) (int, error) {
	// Create a custom client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Apply custom headers
	for _, opt := range opts {
		opt(req)
	}

	// Make the request
	resp, err := client.Do(req)
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