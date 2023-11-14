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

	var request_bytes_buffer []byte = make([]byte, 1024)
	_, err = conn.Read(request_bytes_buffer)
	if err != nil {
		fmt.Println("Error reading from conneciton: ", err.Error())
		os.Exit(1)
	}

	var request_string string = string(request_bytes_buffer)
	lines := strings.Split(request_string, CRLF)

	request_line := strings.Split(lines[0], " ")
	// http_verb := request_line[0]
	path := request_line[1]

	if path == "/"{
		conn.Write([]byte(HTTP_200_OK + CRLF + CRLF))
	}else {
		conn.Write([]byte(HTTP_404_NOT_FOUND + CRLF + CRLF))
	}
}
