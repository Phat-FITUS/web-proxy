package Server

import (
	"github.com/Phat-FITUS/web-proxy/Proxy"
	"fmt"
	"github.com/Phat-FITUS/web-proxy/HTTP"
	"net"
)

func HandleRequest(connection net.Conn){
	header, error := HTTP.GetHeader(connection)

	if (error != nil) {
		fmt.Println("Error" + error.Error())
		connection.Write([]byte(error.Error()))
		connection.Close()
		return
	}

	error = HTTP.ValidateMethod(header)
	if (error != nil) {
		fmt.Println("Error: Method not allow")
		connection.Write([]byte("Method not allow"))
		connection.Close()
		return
	}

	result := HTTP.RedirectRequest(header)
	if (result == "") {
		connection.Write([]byte("Empty Header"))
		connection.Close()
		return
	}

	fmt.Println(result)

	response, error := Proxy.SendRequest(result)

	if (error != nil) {
		connection.Write([]byte(error.Error()))
		connection.Close()
		return
	}

	connection.Write([]byte(response))

	connection.Close()

	fmt.Println("Connection closed!")
}