package main

import (
	"fmt"
	"time"

	. "newknowlage/welcome/core"
)

func main() {
	// Drawwww()
	err := InitConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Drawwww([]string{"热烈欢迎,300,100,0", "中铁建大桥局西南公司,250,20,50", "徐润泽总经理一行,250,50,30", "莅临隧唐基建通指导,300,20,50"})
	Drawwww()
	time.Sleep(3 * time.Second)

}
