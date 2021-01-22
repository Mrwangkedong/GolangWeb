package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

/*
https://tieba.baidu.com/f?kw=%D6%A3%CB%AC&fr=ala0&loc=rec
或者：https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=0
https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=50
https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=100
*/

func main() {
	//指定爬取的起始页、终止页
	var start, end int
	fmt.Println("请输入起始页...")
	_, _ = fmt.Scan(&start)
	fmt.Println("请输入终止页")
	_, _ = fmt.Scan(&end)

	Spider(start, end)
}

func Spider(start int, end int) {
	fmt.Printf("正在爬取第%d页到%d页....\n", start, end)
	//爬取每一页的数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := httpGet(url)
		if err != nil {
			fmt.Println("获取http内容错误：err ", err)
		}
		//fmt.Println(result)
		//将保存的网页内容存到一个文件中去
		f, err := os.Create("G:\\GolangFileText\\爬虫第" + strconv.Itoa(i) + "页.html")
		if err != nil {
			fmt.Println("os.Create：err ", err)
		}
		_, _ = f.Write([]byte(result))
		fmt.Printf("第%d页成功\n", i)
		defer f.Close() //保存好一个文件，关闭一个文件
	}
}

func httpGet(url string) (result string, err error) {
	response, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer response.Body.Close()

	//循环读取网页数据，传出给调用者
	buf := make([]byte, 4096)
	result = ""
	for true {
		n, err2 := response.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次循环
		result += string(buf[:n])
	}

	return result, err

}
