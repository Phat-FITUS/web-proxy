package Proxy


import (
    "errors"
    "strings"
    "fmt"
    "net"
)
const BUFFERSIZE = 1024
func GetHostName(header string) (string, error) {
    temp := strings.Split(header, "\r\n")
    //fmt.Println(temp[1])


    parts := strings.Split(temp[1], ": ")
   
    if (len(parts) == 0) {
        return "", errors.New("Missing HostName")
    }


    return parts[1], nil
}


func SendRequest(header string) {


    hostName, _ := GetHostName(header)
    fmt.Println(hostName)
    conn, err := net.Dial("tcp", hostName+":80")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    fmt.Println("Connected to server.")
 
    _, err = conn.Write([]byte(header))
    if err != nil {
        fmt.Println(err)
        return
    }
    //
    var response string
    buffer := make([]byte, BUFFERSIZE)
    for {
        bytesRead, err := conn.Read(buffer)
        if err != nil {
            break
        }
        response += string(buffer[:bytesRead])
    }
    //conn.Write([]byte(response))
    fmt.Println(response)
}
