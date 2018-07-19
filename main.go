package main

import(
	"fmt"
	"newknowlage/xsql"
	"newknowlage/test"
)



func main() {
	fmt.Println("Start...")
	//snake.SnakeInit()
	//snake.Snake()
	//fmt.Println("myuser&#39; or &#39;foo&#39; = &#39;foo&#39; #")
	//fmt.Println(template.HTMLEscapeString("myuser' or 'foo' = 'foo' #"))
	//fmt.Println(strconv.Quote("myuser&#39; or &#39;foo&#39; = &#39;foo&#39; #"))
	//fmt.Println(strconv.QuoteToGraphic("myuser&#39; or &#39;foo&#39; = &#39;foo&#39; #"))
	//fmt.Println(template.HTMLEscapeString("-++- --/;**/ !!~~~~!"))
	//fmt.Println(html.EscapeString("myuser' or 'foo' = 'foo' #"+"-++- --/'**/ !!~~~~!"))
	//fmt.Println(InjectionDefense.SQLandXSSencode("myuser' or 'foo' = 'foo' #"+"-++- --/'**/ !!~~~~!"))
	//fmt.Println(InjectionDefense.SQLandXSSdecode("myuser&#39; or &#39;foo&#39; = &#39;foo&#39; #-++- --/&#39;**/ !!~~~~!"))

	test.MYsqlTest()

	s := "<IMG src=JaVaScRiPt:alert('XSS')>"
	//s := "\u003c\u0073\u0063\u0072\u0069\u0070\u0074\u003e\u0061\u006c\u0065\u0072\u0074\u0028\u0022\u0078\u0073\u0073\u0022\u0029\u003c\u002f\u0073\u0063\u0072\u0069\u0070\u0074\u003e'"

	fmt.Println(s)
	fmt.Println(xsql.EncodeXSS(s))
	//
	//grpcAndProtoc.GRPCCLIENT()
	//grpcAndProtoc.GRPCSERVER()

	//test.MyMap()
	//var cmd string
	//for{
	//	fmt.Println("downurl-or-CMD:")
	//	fmt.Scan(&cmd)
	//	if cmd == ".quit"{
	//		return
	//	}else{
	//		download.DownLoadInit(cmd)
	//	}
	//}
	//download.DownLoadInit("http://mydmplus.com/res/Mydm/Mydm20180623.zip")
	//download.DownLoadTest()



}

