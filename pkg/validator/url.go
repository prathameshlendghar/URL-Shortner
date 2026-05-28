package customValidator

import (
	"net/url"
	"strings"
)

func IsValidURL(input string) bool {
	// ParseRequestURI is stricter than Parse, it requires a base URI
	u, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	// Force lower case for comparison
	scheme := strings.ToLower(u.Scheme)

	// Block FTP, File, or malicious Javascript schemes
	if scheme != "http" && scheme != "https" {
		return false
	}

	// Ensure there is an actual domain
	if u.Host == "" {
		return false
	}

	return true
}
