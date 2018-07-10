package main

import (
	"newknowlage/mysql"
	"fmt"
	"reflect"
)

func main() {

	//snake.SnakeInit()
	//snake.Snake()


	xs := mysql.InitSql("root","853051095","127.0.0.1","3306","mytest")
	s := mysql.CreateInstance(xs)


	l1 := make(map[string]string)
	l1["del"]="int"
	l1["get_uid"]="int"
	l1["add_uid"]="int"
	l1["remindnum"]="int"
	l1["mobile"]="int"
	l1["distanceTime"]="int"
	l1["get_name"]="string"
	l1["startTipTime"]="string"
	l1["message"]="string"
	l1["wxname"]="string"
	l1["type"]="string"
	l1["intervalTime"]="string"
	l1["unit"]="string"
	l1["add_name"]="string"
	l1["nextTipTime"]="string"

	s.SetTableColType(l1)


	s.Qurey("select * from tipmessage where del=0")
	list := s.Execute()
	for i,v := range list{
		fmt.Println(i,"-",v)
		for i1,v1 := range v{
			fmt.Println(i1,reflect.ValueOf(v1).Type(),"=",v1)
		}
	}
	s.Select("tipmessage","add_uid","int","startTipTime","string","mid","int")
	l2 := make(map[string]interface{})
	l2["del"]=0
	l2["get_name"]="费佳丽"
	s.Where(l2)
	s.AddSuf(" order by mid desc ")
	list = s.Execute()
	for i,v := range list{
		fmt.Println(i,"-",v)
		for i1,v1 := range v{
			fmt.Println(i1,reflect.ValueOf(v1).Type(),"=",v1)
		}
	}

	//l3 := make(map[string]interface{})
	//l3["remindnum"]="100"
	//l3["mobile"]="18398618916"
	//l3["distanceTime"]="3"
	//l3["startTipTime"]="2017-02-23 9:00:00"
	//l3["message"]="回家拖地，"
	//l3["wxname"]="wxid_phg0hogt60io22"
	//l3["type"]="云服务器续费"
	//l3["intervalTime"]="1小时1次，一共3次"
	//l3["unit"]="周"
	//s.Insert(l3,"tipmessage")
	//n := s.ExecuteForLastInsertId()
	//fmt.Println(n)
	//
	s.Update(l2,"tipmessage")
	l4 := make(map[string]interface{})
	l4["mid"]=103
	s.Where(l4)
	n := s.ExecuteForLastInsertId()
	fmt.Println(n)




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
