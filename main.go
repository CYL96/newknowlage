package main

import (
	"newknowlage/Mypprof"
)

func main() {

	//snake.SnakeInit()
	//snake.Snake()
	Mypprof.Mypprofstart()
	defer Mypprof.MypprofStop()
	//MMap1()
	MMap2()

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
func MMap1(){

	for i := 0; i < 100; i++ {
		buf := make(map[int]string, 10240)
		for i := 0; i < 10240; i++ {
			buf[i*10] = "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq"
		}
	}
}

func MMap2(){

	for i := 0; i < 100; i++ {
		buf1 := make(map[int]string)
		for i := 0; i < 10240; i++ {
			buf1[i*10+1] = "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq"
		}
	}
}
