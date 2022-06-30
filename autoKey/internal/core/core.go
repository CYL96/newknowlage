package core

import (
	"fmt"

	"github.com/go-vgo/robotgo"

	. "newknowlage/autoKey/internal/config"
)

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
			if key == CONFIG_CHANGE {
				robotgo.EventEnd()
				go Register()
			}
		}
	}

}
