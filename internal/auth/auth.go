package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(h http.Header) (apiKey string, err error) {
	val := h.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("bad auth")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("bad auth")
	}
	return vals[1], nil
}
