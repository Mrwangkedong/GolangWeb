package main

import (
	"fmt"
	"net/http"
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

	for i := start; i < end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E9%83%91%E7%88%BD&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		httpGet(url)
	}
}

func httpGet(url string) (result string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get(url) err: ", err)
	}

}
