package main

import (
	"fmt"
	"net/http"
)

func main() {
	//注册回调函数，该回调函数会在服务器被访问时，自动调动
	http.HandleFunc("/itcast", handler) //函数的地址就是函数名
	//绑定服务器监听地址
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe() err: ", err)
		return
	}

}

//回调函数，本质：函数指针
func handler(w http.ResponseWriter, r *http.Request) {
	//w 写回给客户端(浏览器)的数据
	//r 从客户端读到的数据
	_, err := w.Write([]byte("Hello World!!"))
	if err != nil {
		fmt.Println("http.ResponseWriter() err: ", err)
		return
	}

}
