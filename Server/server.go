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

	result:= HTTP.RedirectRequest(header)
	if (result == "") {
		connection.Close()
		return
	}

	fmt.Println(result)
	connection.Close()

	fmt.Println("Connection closed!")
}