package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"

	. "newknowlage/autoKey/internal/config"
)

// Register 注册按键绑定函数入口
func Register() {
	registerKey()
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

// registerKey 获取配置文件。将配置的按键进行绑定
func registerKey() {
	config := GetConfig()
	for _, ext := range config.HotKey {
		keys := strings.Split(ext.Key, ",")
		fmt.Println("注册:", ext.Name, keys, ":", ext.Account, ext.Pwd)
		robotgo.EventHook(hook.KeyDown, keys, getFunc(ext.Name, ext.Account, ext.Pwd))
	}
}

// getFunc 获取对应热键绑定的执行函数
func getFunc(name, account, pwd string) func(hook.Event) {
	return func(e hook.Event) {
		fmt.Println("识别:", name)
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
