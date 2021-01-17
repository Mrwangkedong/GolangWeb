package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	MyListener, err := net.Listen("tcp", "127.0.0.1:8000")
	//处理错误
	errFunc(err, "net.listen()")
	//if err != nil{
	//	fmt.Println("net.listen err:",err)
	//	return
	//}

	conn, err := MyListener.Accept()
	errFunc(err, "MyListener.Accept()")
	defer MyListener.Close()
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	errFunc(err, "conn.Read(buf)")
	fmt.Println(string(buf[:n]))

}

//错误处理函数封装
func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, "  err:", err)
		os.Exit(1) //将当前--进程--结束！
	}
}
