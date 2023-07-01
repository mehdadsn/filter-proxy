package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting Server...")

	// listen on port 9000
	server, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Error Listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Server Started, Listening on Port 9000 ....")

	// accept connection

	// run loop forever (or until ctrl-c)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		//fmt.Println("Client Connected\n")
		go processConnection(connection)
		// get message, output

	}
}

func processConnection(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	headers := string(buffer[:mLen])
	//fmt.Println(headers)
	scanner := bufio.NewScanner(strings.NewReader(headers))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Host:") {
			line = line[strings.Index(line, " ")+1 : len(line)]
			fmt.Println(line)
		}
	}
	_, err = connection.Write([]byte(string(buffer[:mLen])))
	connection.Close()
}
