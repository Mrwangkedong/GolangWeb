package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
练习：
	访问某一文件夹下面的图片，音频等文件
*/

func myHandle(w http.ResponseWriter, r *http.Request) {
	//判断要访问的文件是否存在
	flag := 0 //表示不存在
	//获取客户端要访问的文件名
	fileName := r.URL.String()[1:]
	fmt.Println("fileName:", fileName)
	//打开默认文件夹
	catalog, err := os.OpenFile("G://GolangFileText", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("os.OpenFile(，，os.ModeDir) err: ", err)
		return
	}
	//获取文件夹下的所有文件
	files, err := catalog.Readdir(-1) //<=0，返回All，>=0，返回指定数目
	for _, item := range files {
		fmt.Printf("%s\t", item.Name())
		if item.Name() == fileName {
			flag = 1 //表示存在
			break    //退出当前循环
		}
	}

	if flag == 1 {
		WriteToClient(w, fileName) //将用户想要访问的文件发送过去
	} else {
		_, _ = w.Write([]byte("访问文件不存在..."))
	}

}

func WriteToClient(w http.ResponseWriter, fileName string) {
	freader, err := os.OpenFile("G://GolangFileText/"+fileName, os.O_RDWR, 6)
	defer freader.Close()
	if err != nil {
		fmt.Println("os.OpenFile err: ", err)
		return
	}
	//3.创建buffer,读文件并写入client
	buf := make([]byte, 1024*4)
	for {
		n, err := freader.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完了、、、、")
			return
		}
		_, _ = w.Write(buf[:n])
	}

}

func main() {
	http.HandleFunc("/", myHandle) //函数的地址就是函数名
	//绑定服务器监听地址
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe() err: ", err)
		return
	}
}
