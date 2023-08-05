package Proxy

import (
	"errors"
	"fmt"
	"net"
	"strings"
)
const BUFFER_SIZE = 1024

func GetHostName(header string) (string, error) {
    temp := strings.Split(header, "\r\n")

    parts := strings.Split(temp[1], ": ")

    if (len(parts) == 0) {
        return "", errors.New("missing HostName")
    }

    return parts[1], nil
}


func SendRequest(header string) (string, error) {
    hostName, _ := GetHostName(header)

    conn, err := net.Dial("tcp", hostName+":80")
    if err != nil {
        fmt.Println(err)
        return "", errors.New(err.Error())
    }
    defer conn.Close()
    fmt.Println("Connected to", hostName)

    _, err = conn.Write([]byte(header))
    if err != nil {
        fmt.Println(err)
        return "", errors.New(err.Error())
    }

    var response string
    buffer := make([]byte, BUFFER_SIZE)
    for {
        bytesRead, err := conn.Read(buffer)
        if err != nil {
            break
        }
        response += string(buffer[:bytesRead])
    }

    fmt.Println(response)
    return response, nil
}
