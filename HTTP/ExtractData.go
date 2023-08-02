package HTTP

import (
	"errors"
	"strings"
)

func GetMethod(header string) (string, error) {
	request := header[ : strings.Index(header, "\r\n")]
	parts := header[ : strings.Index(request, " /")]

	if (len(parts) == 0) {
		return "", errors.New("missing method")
	}

	return parts, nil
}

func GetRequest(header string) (string) {
	request := header[ : strings.Index(header, "\r\n")]
	toDelete, _ := GetURL(header)

	return strings.ReplaceAll(request, toDelete, "")
}

func GetURL(header string) (string, error) {
	url := header[strings.Index(header, "/") + 1 : strings.Index(header, " HTTP")]

	if (len(url) == 0) {
		return "", errors.New("missing url")
	}

	return url, nil
}
