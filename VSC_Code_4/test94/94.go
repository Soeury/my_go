package main

import (
	"fmt"
	"regexp"
)

/*

    模板如下 ：

	    str := "需要查找的字符串"

	    reg := regexp.MustCompile(`正则表达式`)
	    if reg == nil {
		    fmt.Println("Found err!")
	    }

	    ret := reg.FindAllString(str , -1)  或者  ret := reg.FindAllStringSubMatch(str , -1)
	    fmt.Println(ret)

*/

func main() {

	// 正则表达式（能使用 string 和 strconv 就不使用正则表达式 ，正则表达式效率低，代码可读性弱）
	// 不用背下来 ， 要用的时候去查就可以

	str1 := "abc arc a9c tac rfc 9a0 a84 auc"
	str2 := "3.14 89.45 785 1. 0.2 456.123 119"

	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("Found err!")
		return
	}

	reg2 := regexp.MustCompile(`(\d+\.\d+)`)
	if reg2 == nil {
		fmt.Println("Found err!")
		return
	}

	ret1 := reg1.FindAllStringSubmatch(str1, -1)
	fmt.Println(ret1)
	ret2 := reg2.FindAllString(str2, -1)
	fmt.Println(ret2)

	// 正则表达式提取 html 中需要的部分
	str3 := `
	
<!DOCTYPE html>
<html class="">
<head>
<meta name="wechat-enable-text-zoom-em" content="true">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="color-scheme" content="light dark">
<meta name="viewport" content="width=device-width,initial-scale=1.0,maximum-scale=1.0,user-scalable=0,viewport-fit=cover">
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-status-bar-style" content="black">
<meta name="format-detection" content="telephone=no">
    <div>love</div>
	<div>forever</div>
	<div>you
	are
	my 
	only
	</div>
	<div>234</div> 
<meta name="referrer" content="origin-when-cross-origin">
<meta name="referrer" content="strict-origin-when-cross-origin">
<script type="text/javascript" nonce="1752922497" reportloaderror>
  window.logs = { pagetime: {} };
  window.logs.pagetime['html_begin'] = (+new Date());
  window.LANG = "zh_CN";
</script>
</head>
	
	`

	reg3 := regexp.MustCompile(`<div>((.*))</div>`)
	if reg3 == nil {
		fmt.Println("Found err!")
		return
	}

	ret3 := reg3.FindAllStringSubmatch(str3, -1)

	// 不带<div></div>
	for _, ret := range ret3 {
		fmt.Println(ret[0])
	}

	// 带<div></div>
	for _, ret := range ret3 {
		fmt.Println(ret[1])
	}

}
