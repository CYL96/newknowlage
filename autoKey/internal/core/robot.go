package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"

	. "newknowlage/autoKey/internal/config"
)

func Register() {
	registerKey()
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func registerKey() {
	config := GetConfig()
	for _, ext := range config.HotKey {
		keys := strings.Split(ext.Key, ",")
		fmt.Println("注册:", ext.Name, keys, ":", ext.Account, ext.Pwd)
		robotgo.EventHook(hook.KeyDown, keys, getFunc(ext.Account, ext.Pwd))
	}
}

func getFunc(account, pwd string) func(hook.Event) {
	return func(e hook.Event) {
		fmt.Println("识别")
		time.Sleep(150 * time.Millisecond)
		robotgo.TypeStr(account)
		robotgo.KeyDown("tab")
		time.Sleep(50 * time.Millisecond)
		robotgo.KeyUp("tab")

		robotgo.TypeStr(pwd)
		robotgo.KeyDown("enter")
		time.Sleep(50 * time.Millisecond)
		robotgo.KeyUp("enter")
	}
}
