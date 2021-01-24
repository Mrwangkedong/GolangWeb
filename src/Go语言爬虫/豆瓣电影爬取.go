package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

/*
https://movie.douban.com/top250
https://movie.douban.com/top250?start=25&filter=
https://movie.douban.com/top250?start=50&filter=
*/
/*
思路
获取html函数
对单页HTML进行爬取
*/
//电影结构体定义
type movie struct {
	Chinese_Name string  //中文名
	English_Name string  //英文名
	Other_Name   string  //香港、台湾翻译电影名
	Director     string  //导演名
	Contray      string  //国家
	Time         string  //上映时间
	MovieType    string  //电影类型
	Grade        float32 //评分
}

func main() {

	for i := 0; i < 1; i++ {
		SpiderOnePage(i)
	}
}

func SpiderOnePage(pageNum int) {
	//1.获取网页源码
	result, err := httpGet_Douban("http://movie.douban.com/top250?start=" + strconv.Itoa((pageNum)*25) + "&filter=")
	fmt.Println("http://movie.douban.com/top250?start=" + strconv.Itoa((pageNum)*25) + "&filter=")
	if err != nil {
		fmt.Println("httpGet_Douban() err: ", err)
		return
	}

	//2.进行解析
	ChineseName := regexp.MustCompile("<div class=\"hd\">.*\\n.*\\n.*<span class=\"title\">(?s:(.*?))</span>").FindAllStringSubmatch(result, -1)
	for _, item := range ChineseName {
		fmt.Println(item[1])
	}

}

//获得网页源码
func httpGet_Douban(url string) (result string, err error) {
	response, err1 := http.Get(url)
	fmt.Println(url)
	fmt.Println(response.Body)
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
