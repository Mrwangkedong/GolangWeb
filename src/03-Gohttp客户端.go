package main

import (
	"fmt"
	"net/http"
)

func main() {
	//获取服务器，应答包
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("http.Get() err: ", err)
		return
	}
	defer resp.Body.Close()

	//简单查看应答包.
	fmt.Println("Header: ", resp.Header)
	fmt.Println("Status", resp.Status)
	fmt.Println("Proto", resp.Proto)
	fmt.Println("Body", resp.Body)

	buf := make([]byte, 4096)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println(123)
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		result += string(buf[:n])
	}
	fmt.Println(result)

}
