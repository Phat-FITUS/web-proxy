package HTTP

import (
	"fmt"
	"strings"
)

func Mapify(input string, separate string) (map[string]string) {
	temp := input[:strings.Index(input, separate)]
	input = input[:0] + input[len(temp):]
	result := make(map[string]string)


	// Divide input into part key-value
	pairs := strings.Split(input, separate)
	for i:=0; i < len(pairs); i++ {
		pair:= pairs[i]
		parts := strings.SplitN(pair, ":", 2)
		if (len(parts) == 2) {
			// Trimspace delete whitespace
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			result[key] = value
		}
	}
	return result
}

func CreateDirectRequest(result map[string]string) string {
	requestContent := fmt.Sprintf("Host: %s\r\n", result["Host"])
	// Add other headers to the request content
	for key, value := range result {
		if key != "Host" {
			requestContent += fmt.Sprintf("%s: %s\r\n", key, value)
		}
	}
	// Add an empty line to mark the end of the headers
	requestContent += "\r\n"
	return requestContent
}
