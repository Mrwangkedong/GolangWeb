package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "abc a7c mfc cat"
	//1.解析，编译正则表达式
	//re := regexp.MustCompilePOSIX(".{2}c")
	re := regexp.MustCompilePOSIX("a[0-9]c")
	//2.提取所需信息
	strList := re.FindAllStringSubmatch(str, -1)

	fmt.Println(strList)

	//匹配小数
	strr := "3.14 123.123 .68 hahha 1.0 abc 3.66 123"
	//1.解析，编译正则表达式
	ret := regexp.MustCompile("[0-9]+\\.[0-9]+") //[0-9]匹配数字  + 1~N次 \\.匹配.
	//2.提取所需信息
	strList = ret.FindAllStringSubmatch(strr, -1)
	fmt.Println(strList)

	//go语言定义多行字符串，利用 ` `
	str = `<div class="hd">
                        <a href="https://movie.douban.com/subject/1308857/" class="">
                            <span class="title">可可西里</span>
                                <span class="other">&nbsp;/&nbsp;Kekexili: Mountain Patrol</span>
                        </a>`
	//提取标签<frame>标签中的src
	//1.解析编译正则表达式
	re = regexp.MustCompile(`<div class="hd">.*\n.*\n.*<span class="title">(?s:(.*?))</span>`)
	//2.获得结果
	strList = re.FindAllStringSubmatch(str, -1)
	fmt.Println(strList)

	//for _,item := range strList{
	//	fmt.Printf("item[0]: %s\titem[1]: %s\n",item[0],item[1])
	//}

	/*输出
	[[<frame src="/static
	/pkgdoc/i.html" /static
	/pkgdoc/i.html] [<frame name="main" src="/static/pkgdoc/main.html" /static/pkgdoc/main.html]]
	*/

}
