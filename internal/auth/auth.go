package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("first part of header is malformed")
	}

	return vals[1], nil
}
