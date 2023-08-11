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

    header, body, err := HTTP.GetResponse(conn)

    return header + body, err
}
