package core

import (
	"fmt"

	"github.com/go-vgo/robotgo"

	. "newknowlage/autoKey/internal/config"
)

// Run 运行函数入口
func Run() {
	InitConfigCh()
	err := InitConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go CheckConfig()
	go Register()
	for {
		select {
		case key := <-ConfigCh:
			switch key {
			case CONFIG_CHANGE:
				robotgo.EventEnd()
				go Register()
			case CONFIG_ERROR:
				robotgo.EventEnd()
				return
			}
		}
	}

}
