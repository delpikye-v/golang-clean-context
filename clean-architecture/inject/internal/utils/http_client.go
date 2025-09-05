package utils

import "net/http"

// NewHTTPClient returns a shared HTTP client
func NewHTTPClient() *http.Client {
	return &http.Client{}
}
