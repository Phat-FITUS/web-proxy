package Server

import (
	"fmt"
	"net"
	"github.com/Phat-FITUS/web-proxy/HTTP"
)

func HandleRequest(connection net.Conn){
	header, error := HTTP.GetHeader(connection)

	if (error != nil) {
		fmt.Println("Error" + error.Error())
		connection.Close()
		return
	}

	fmt.Println("Received request:", header)

	connection.Close()

	fmt.Println("Connection closed!")
}