package HTTP

import (
	"errors"
	"strings"
	"golang.org/x/exp/slices"
)

func ValidateHeader(header string) error {
	var err error
	_, err = GetMethod(header)
	if err != nil {
		return err
	}

	_, err = GetURL(header)
	if err != nil {
		return err
	}

	return nil
}

var ALLOW_METHODS = []string{"GET", "POST", "HEAD"}

func ValidateMethod(header string) error {
	method, err := GetMethod(header)

	if (!slices.Contains(ALLOW_METHODS, method) || err != nil) {
		return errors.New("invalid method")
	}

	return nil
}

func IsMediaFetching(header string) bool {
	h := Mapify(header, "\r\n")
	acc, exist := h["Accept"]
	if (exist) {
		return !strings.Contains(acc, "text")
	}
	return false
}