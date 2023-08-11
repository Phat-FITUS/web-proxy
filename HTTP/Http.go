package HTTP

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	LF = 10
	CR = 13
)

//End of a header of HTTP Request
var EndMark = [4]byte{CR, LF, CR, LF}

const BUFFER_SIZE = 1024

//Get header of request
func GetHeader(con net.Conn) (string, error){
	header := ""
	buffer := make([]byte, BUFFER_SIZE)

	for {
		n, err := con.Read(buffer)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			fmt.Println("Error reading response:", err)
			return "", err
		}

		header += string(buffer[:n])

		if (strings.Contains(header, "\r\n\r\n")){
			break
		}
	}

	return header, nil
}

func GetBody(con net.Conn, header string, prevOverflow string) string {
	h := Mapify(header, "\r\n")
	length, isExist := h["Content-Length"]
	body := prevOverflow

	if (isExist) {
		length, _ := strconv.Atoi(length)
		currentByte := len(body)

		for currentByte <= length {
			buffer := make([]byte, BUFFER_SIZE)
			n, err := con.Read(buffer)

			if err != nil {
				if err.Error() == "EOF" {
					break
				}

				fmt.Println("Error reading response:", err)
				return ""
			}

			body += string(buffer[:n])
			currentByte += n
		}
	}

	return body
}

func GetResponse(connection net.Conn) (string, string, error) {
	header, error := GetHeader(connection)
	endHeaderPos := strings.Index(header, "\r\n\r\n")

	body := ""
	if (endHeaderPos != -1 && endHeaderPos + 4 < len(header)) {
		body = GetBody(connection, header, header[endHeaderPos + 4:])
		header = header[:endHeaderPos + 4]
	}

	return header, body, error
}

//Redirect the request from this proxy to the destination
func RedirectRequest(request string) (string) {
	err := ValidateHeader(request)

	if (err != nil) {
		fmt.Println(err)
		return ""
	}

	requestContent := GetRequest(request)

	tempMap := Mapify(request, "\r\n")

	tempMap["Connection"] = "close"

	newRequest := fmt.Sprintf("%s \r\n", requestContent)
	newRequest += CreateDirectRequest(tempMap)

	return newRequest
}