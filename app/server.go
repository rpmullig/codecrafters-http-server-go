package main

import (
	"fmt"
	 "net"
	 "os"
	"strings"
	"strconv"
)

const CRLF string = "\r\n"
const HTTP_200_OK string = "HTTP/1.1 200 OK"
const HTTP_404_NOT_FOUND string = "HTTP/1.1 404 Not Found"
const CONTENT_TYPE_TEXT string ="Content-Type: text/plain"
const ECHO_PREFIX string = "/echo/"

func handle(conn net.Conn) {
	defer conn.Close()
	var request_bytes_buffer []byte = make([]byte, 1024)
	_, err := conn.Read(request_bytes_buffer)
	if err != nil {
		fmt.Println("Error reading from conneciton: ", err.Error())
		os.Exit(1)
	}

	var request_string string = string(request_bytes_buffer)
	lines := strings.Split(request_string, CRLF)

	request_line := strings.Split(lines[0], " ")
	// http_verb := request_line[0]
	path := request_line[1]

	request_headers, i := parse_headers(lines)
	request_body := lines[i:]
	fmt.Println("Request body: ", request_body)

	if path == "/" {
		conn.Write([]byte(HTTP_200_OK + CRLF + CRLF))
	} else if strings.HasPrefix(path, ECHO_PREFIX) {
		// response_body, _ := strings.CutPrefix(path, ECHO_PREFIX) available in version 1.20+
		var response_body string = path[len(ECHO_PREFIX):]
		conn.Write([]byte(HTTP_200_OK + CRLF + CONTENT_TYPE_TEXT + CRLF + "Content-Length: " + strconv.Itoa(len(response_body)) + CRLF + CRLF + response_body))
	} else if path == "/user-agent" {
		conn.Write([]byte(HTTP_200_OK + CRLF + CONTENT_TYPE_TEXT + CRLF + "Content-Length: " + strconv.Itoa(len(request_headers["User-Agent"])) + CRLF + CRLF + request_headers["User-Agent"]))
	} else {
		conn.Write([]byte(HTTP_404_NOT_FOUND + CRLF + CRLF))
	}
}

func parse_headers(request_lines []string) (map[string]string, int) {
    headers := make(map[string]string)
    var i int
    for i = 1; i < len(request_lines) && request_lines[i] != ""; i++ {
        split_line := strings.SplitN(request_lines[i], ":", 2)
        if len(split_line) == 2 {
            key := strings.TrimSpace(split_line[0])
            value := strings.TrimSpace(split_line[1])
            headers[key] = value
        }
    }
    return headers, i
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle(conn)
	}

}


