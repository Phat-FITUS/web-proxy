package main

import (
	"fmt"
	"os"
	"github.com/Phat-FITUS/web-proxy/Server"
	"net"
);

func main() {
	server, err := net.Listen("tcp", "localhost:8080")

    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

	fmt.Println("Listening at port 8080")

    defer server.Close()

    for {
        connection, err := server.Accept()

        if err != nil {
            fmt.Println("Error: ", err.Error())
            os.Exit(1)
        }

        fmt.Println("New Connection")

        go Server.HandleRequest(connection)
    }
}