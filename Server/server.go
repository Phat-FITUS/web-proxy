package Server

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"github.com/Phat-FITUS/web-proxy/HTTP"
	"github.com/Phat-FITUS/web-proxy/Proxy"
)

func HandleRequest(connection net.Conn){
	if (!IsInService()) {
		content, _ := os.ReadFile("./Server/status/exceed.html")
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 403 Forbidden\r\nContent-type: text/html\r\nContent-length: %d\r\n\r\n%s", len(string(content)), string(content))))
		connection.Close()
		return
	}

	header, body, error := HTTP.GetResponse(connection)

	if (error != nil) {
		content, _ := os.ReadFile("./Server/status/403.html")
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 403 Forbidden\r\nContent-type: text/html\r\nContent-length: %d\r\n\r\n%s", len(string(content)), string(content))))
		connection.Close()
		return
	}

	error = HTTP.ValidateMethod(header)
	if (error != nil) {
		content, _ := os.ReadFile("./Server/status/403.html")
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 403 Forbidden\r\nContent-type: text/html\r\nContent-length: %d\r\n\r\n%s", len(string(content)), string(content))))
		connection.Close()
		return
	}

	host, _ := HTTP.GetHostName(header)
	if (!IsAcceptableHost(host)) {
		content, _ := os.ReadFile("./Server/status/403.html")
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 403 Forbidden\r\nContent-type: text/html\r\nContent-length: %d\r\n\r\n%s", len(string(content)), string(content))))
		connection.Close()
		return
	}

	url := header[strings.Index(header, " ") + 1 : strings.Index(header, " HTTP")]
	url = strings.ReplaceAll(url, "http://", "")

	file := url[strings.LastIndex(url, "/"):]
	requestedFile := "./cache/" + url

	if (HTTP.IsMediaFetching(header)) {
		data, error := os.ReadFile(requestedFile)
		fileInfo, _ := os.Stat(requestedFile)

		if (error == nil && time.Since(fileInfo.ModTime()).Seconds() <= GetConfigTime()) {
			connection.Write([]byte(string(data)))
			connection.Close()
			fmt.Println("Cache Returning")
			return
		}
	}

	redirectedHeader := HTTP.RedirectRequest(header)
	if (redirectedHeader == "") {
		content, _ := os.ReadFile("./Server/status/403.html")
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 403 Forbidden\r\nContent-type: text/html\r\nContent-length: %d\r\n\r\n%s", len(string(content)), string(content))))
		connection.Close()
		return
	}
	fmt.Printf("Received:\n%s", redirectedHeader + body)

	response, error := Proxy.SendRequest(redirectedHeader + body)

	if (error != nil) {
		content, _ := os.ReadFile("./Server/status/404.html")
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 404 Not Found!\r\nContent-type: text/html\r\nContent-length: %d\r\n\r\n%s", len(string(content)), string(content))))
		connection.Close()
		return
	}

	if (HTTP.IsMediaFetching(header)) {
		pathToFile :=  strings.TrimSuffix(requestedFile, file)
		os.MkdirAll(pathToFile, os.ModePerm)
		os.WriteFile(requestedFile, []byte(response), 0644)
	}

	fmt.Println("Response returned")
	connection.Write([]byte(response))

	connection.Close()

	fmt.Println("Connection closed!")
}