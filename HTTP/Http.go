package HTTP

import (
	"bytes"
	"fmt"
	"net"
)

const (
	LF = 10
	CR = 13
)

//End of a header of HTTP Request
var EndMark = [4]byte{CR, LF, CR, LF}

//Get header of request
func GetHeader(con net.Conn) (string, error){
	header := ""
	buffer := make([]byte, 1024)

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

		if (bytes.Equal(buffer[n - 4 : n], EndMark[:])){
			break
		}
	}

	return header, nil
}

//Redirect the request from this proxy to the destination
func RedirectRequest(request string) string {
	return ""
}