package main

import (
	"fmt"
	"net"
)

//装作浏览器
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial() err: ", err)
		return
	}
	defer conn.Close()

	httpRequest := "GET /itcast HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"

	conn.Write([]byte(httpRequest))

	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))

	/*
		HTTP/1.1 200 OK
		Date: Mon, 18 Jan 2021 02:33:32 GMT
		Content-Length: 13
		Content-Type: text/plain; charset=utf-8

		Hello World!!

		Process finished with exit code 0

	*/
}
