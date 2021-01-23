package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
https://tieba.baidu.com/f?kw=%D6%A3%CB%AC&fr=ala0&loc=rec
或者：https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=0
https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=50
https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=100
*/

/*
1.提示用户指定爬的页数
2.使用start，end爬取每一页
3.获取每一页url ———— 下一页 = 前一页+50
4.封装实现HttpGet()，爬取一个网页的数据内容，通过result返回
5.创建 .html文件。使用循环因子i命名。
6.将result写入文件，使用f.close()关闭，不推荐使用defer
*/

//创建全局channel
var ch = make(chan int)

func main() {
	//指定爬取的起始页、终止页
	var start, end int
	fmt.Println("请输入起始页...")
	_, _ = fmt.Scan(&start)
	fmt.Println("请输入终止页")
	_, _ = fmt.Scan(&end)

	ch = make(chan int, end-start+1)

	Spider2(start, end)

}

//爬取单个页面的函数
func SpiderPage(url string, i int, end int) {
	result, err := httpGet2(url)
	if err != nil {
		fmt.Println("获取http内容错误：err ", err)
		return
	}
	//fmt.Println(result)
	//将保存的网页内容存到一个文件中去
	f, err := os.Create("G:\\GolangFileText\\爬虫第" + strconv.Itoa(i) + "页.html")
	if err != nil {
		fmt.Println("os.Create：err ", err)
		return
	}
	_, _ = f.Write([]byte(result))
	fmt.Printf("第%d页成功\n", i)
	f.Close() //保存好一个文件，关闭一个文件

	ch <- i

}

//进行爬取
func Spider2(start int, end int) {
	fmt.Printf("正在爬取第%d页到%d页....\n", start, end)
	//开始时间
	nowtime := time.Now()
	fmt.Println("开始时间：", nowtime)
	//循环爬取
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		//爬取每一页的数据
		go SpiderPage(url, i, end)
	}
	//主go程
	for true {
		if len(ch) == end-start+1 {
			//结束时间
			fmt.Println("结束时间：", time.Now())
			return
		}
	}
}

//获得网页源码
func httpGet2(url string) (result string, err error) {
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
