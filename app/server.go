package main

import (
	"fmt"
	 "net"
	 "os"
	 "strings"
)

func main() {
	const CRLF string = "\r\n"
	const HTTP_200_OK string = "HTTP/1.1 200 OK"
	const HTTP_404_NOT_FOUND string = "HTTP/1.1 404 Not Found"

	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	var request_bytes []byte
	_, err = conn.Read(request_bytes)
	if err != nil {
		fmt.Println("Error reading from conneciton: ", err.Error())
		os.Exit(1)
	}

	var request_string string = strings.Reader().Read(request_bytes)
	fmt.Println(request_string)

	conn.Write([]byte(HTTP_200_OK + CRLF + CRLF))
}
