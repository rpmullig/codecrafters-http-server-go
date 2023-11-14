package main

import (
	"fmt"
	 "net"
	 "os"
)

func main() {
	const CRLF string = "\r\n"
	const HTTP_200_OK string = "HTTP/1.1 200 OK"

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

	defer l.Close()
	conn.Write([]byte(HTTP_200_OK + CRLF + CRLF))
 }
