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
	str = `<html lang="zh-CN">
	<head>
	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	<link rel="shortcut icon" href="/static/img/go.ico">
	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
	</head>
	<frameset cols="15,85">
	<frame src="/static
/pkgdoc/i.html">
	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
	<noframes>
	</noframes>
	</frameset>
	</html>`
	//提取标签<frame>标签中的src
	//1.解析编译正则表达式
	re = regexp.MustCompile(`<frame.*src="(?s:(.*?))"`)
	//2.获得结果
	strList = re.FindAllStringSubmatch(str, -1)
	fmt.Println(strList)

	/*输出
	[[<frame src="/static
	/pkgdoc/i.html" /static
	/pkgdoc/i.html] [<frame name="main" src="/static/pkgdoc/main.html" /static/pkgdoc/main.html]]
	*/

}
