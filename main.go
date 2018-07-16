package main

import(
	"fmt"
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

