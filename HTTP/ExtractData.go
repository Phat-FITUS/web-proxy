package HTTP

import (
	"errors"
	"strings"
)

func GetMethod(header string) (string, error) {
	request := header[ : strings.Index(header, "\r\n")]
	parts := header[ : strings.Index(request, " ")]

	if (len(parts) == 0) {
		return "", errors.New("missing method")
	}

	return parts, nil
}

func GetRequest(header string) (string) {
	request := header[ : strings.Index(header, "\r\n")]
	h, _ := GetHostName(header)

	toDelete := "http://" + h

	return strings.ReplaceAll(request, toDelete, "")
}

func GetHostName(header string) (string, error) {
    temp := strings.Split(header, "\r\n")

    parts := strings.Split(temp[1], ": ")

    if (len(parts) == 0) {
        return "", errors.New("missing HostName")
    }

    return parts[1], nil
}

func GetURL(header string) (string, error) {
	host, _ := GetHostName(header)
	url := strings.ReplaceAll(header[strings.Index(header, " ") + 1 : strings.Index(header, " HTTP")], host, "")

	if (len(url) == 0) {
		return "", errors.New("missing url")
	}

	return url, nil
}
