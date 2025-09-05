package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetJSON sends GET request and decodes JSON into target
func GetJSON(client *http.Client, url string, target any) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed GET %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed read body: %v", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed unmarshal JSON: %v", err)
	}
	return nil
}
