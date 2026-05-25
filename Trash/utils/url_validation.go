package utils

import (
	"fmt"
	"net/url"
)

func ValidateLongUrl(longUrl string) (string, error) {
	parsedUrl, err := url.Parse(longUrl)

	if err != nil {
		errStr := fmt.Errorf("Error: %v", err)
		return "", errStr
	}

	if parsedUrl.Scheme != "" && parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		errStr := fmt.Errorf("Error: Supported Long URL methods are http and https")
		return "", errStr
	}

	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}

	parsedUrl.RawQuery, _ = url.QueryUnescape(parsedUrl.RawQuery)
	parsedLongUrl := parsedUrl.String()
	return parsedLongUrl, nil
}

func ValidateShortUrl(shortUrl string) (string, error) {
	parsedUrl, err := url.Parse(shortUrl)

	if err != nil {
		errStr := fmt.Errorf("Error: %v", err)
		return "", errStr
	}

	if parsedUrl.Scheme != "" && parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		errStr := fmt.Errorf("Error: Supported Long URL methods are http and https")
		return "", errStr
	}
	parsedUrlPath := parsedUrl.Path[1:]
	return parsedUrlPath, nil
}
