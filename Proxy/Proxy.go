package Proxy

import (
	"errors"
	"fmt"
	"net"
	"github.com/Phat-FITUS/web-proxy/HTTP"
)
const BUFFER_SIZE = 1024

func SendRequest(header string) (string, error) {
    hostName, _ := HTTP.GetHostName(header)

    conn, err := net.Dial("tcp", hostName+":80")
    if err != nil {
        fmt.Println(err)
        return "", errors.New(err.Error())
    }
    defer conn.Close()

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

    return response, nil
}
