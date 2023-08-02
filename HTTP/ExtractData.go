package HTTP

import (
	"strings"
	"errors"
)

func GetMethod(header string) (string, error) {
	var err error
	temp:= strings.Split(header, "\n")
	parts:= strings.Split(temp[0], " ")
	if (len(parts[0]) == 0) {
		err = errors.New("Missing method")
		return "", err
	}
	return parts[0], nil
}

func GetRequest(header string) (string) {
	temp := strings.Split(header, "\n")	
	value := temp[0]
	toDelete, _ := GetURL(header)

	index:= strings.Index(value, toDelete)
	newMethod := value[:index] + value[index + len(toDelete) :]

	return newMethod
}

func GetURL(header string) (string, error) {
	var err error
	parts := strings.Split(header, " ")

	if (len(parts[1]) == 0) {
		err = errors.New("Missing method")
		return "", err
	}

	if strings.HasPrefix(parts[1], "/") {
		parts[1] = parts[1][1:]
	}

	return parts[1], nil
}

func CheckRequest(header string) error {
	var err error
	_, err = GetMethod(header)
	if (err != nil) {
		return err
	}

	_, err = GetURL(header)
	if (err != nil) {
		return err
	}
	
	return nil
}