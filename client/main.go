package main

import (
	"fmt"
	"net"
)

const (
	Server_Host = "127.0.0.1"
	Server_Port = "9000"
	Server_Type = "tcp"
)

func main() {
	//establish connection
	connection, err := net.Dial(Server_Type, Server_Host+":"+Server_Port)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	//send some data
	_, err = connection.Write([]byte("Hello from client!"))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))

}
