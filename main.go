package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

const (
	listFile = "list.txt"
)

var sites []string

func main() {
	sites = getSiteList()
	//for _, v := range sites {
	//fmt.Println(v)
	//}
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
	//test, err := http.Get("https://www.google.com")
	//fmt.Println(test)
	if err != nil {
		fmt.Println(err.Error())
	}

	// run loop forever (or until ctrl-c)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go processConnection(connection)
		// get message, output

	}
}

func processConnection(connection net.Conn) {
	fmt.Println("\n######### New Request ##########\n")
	defer connection.Close()
	buffer := make([]byte, 4096)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}

	request := string(buffer[:mLen])
	fmt.Printf(request)
	headers := strings.Split(request, "\n")
	var host string
	var port string
	var url string
	//var forwardRequest string
	for _, v := range headers {
		if len(v) >= 4 {
			if v[:4] == "Host" {
				hostport := strings.Split(v[6:], ":")
				host = hostport[0]
				if len(hostport) > 1 {
					port = hostport[1][:len(hostport[1])-1]
				} else {
					host = host[:len(host)-1]
				}
			}
		}
	}
	fmt.Println("************")
	//fmt.Println(len(host))
	//fmt.Println(host)
	fmt.Println(port)
	fmt.Println(len(port))
	fmt.Println(url)
	fmt.Println("************")
	if checkList(host) {
		res := "HTTP/1.1 403 Forbidden\r\nContent-Type: text/html; charset=utf-8\r\n\r\n<h1>Not Allowed</h1>"
		connection.Write([]byte(res))
	} else {

		fmt.Println(host)
		connection2, err := net.Dial("tcp", host+":80")
		if err != nil {
			fmt.Println("Error dialing Connection2: " + err.Error())
		}
		if connection2 != nil {
			defer connection2.Close()
			_, err = connection2.Write([]byte(request))
			buffer2 := make([]byte, 4096)
			mLen, err := connection2.Read(buffer2)
			fmt.Println("flag$$$$$$$$")

			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(string(buffer2))
			n, err := connection.Write(buffer2[:mLen])
			fmt.Println(mLen, " ", n)
			if err != nil {
				fmt.Println("errpr writing: ", err.Error())
			}
		}

	}
	fmt.Println("\n######### Request End ##########\n")

}

func getSiteList() (sites []string) {
	data, err := ioutil.ReadFile(listFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file read: %v\n", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		sites = append(sites, line)
	}
	return sites
}

func checkList(host string) bool {
	//fmt.Println("_ _ _ _ _ _ _ _check list_ _ _ _")
	//fmt.Println(host)

	for _, v := range sites {
		//fmt.Println(v)
		if v[0] == '*' && v[2] == '*' {
			fmt.Printf("filterd according to %v\n", v)
			return true
		}

		if v[0:3] == "www" && len(host) > len(v[4:]) {
			if host[len(host)-len(v[4:]):] == v[4:] {
				fmt.Printf("filterd according to %v\n", v)
				return true
			}
		}

		if v[0] == '*' && (len(host) > len(v)-2) {
			//fmt.Println(v[2 : len(v)-1])
			//fmt.Println(host[len(host)-len(v[2:len(v)-1]):])
			if v[2:len(v)-1] == host[len(host)-len(v[2:len(v)-1]):] {
				fmt.Printf("filterd according to %v\n", v)
				return true
			}
		}
	}
	//fmt.Println("_ _ _ _ _ _ _ _End check list _ _ _ _")
	return false
}
