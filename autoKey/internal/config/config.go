package config

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var Config *ConfigExt
var configTime int64
var ConfigCh chan RunType

// GetConfig 获取当前配置文件
func GetConfig() *ConfigExt {
	return Config
}

// InitConfigCh 初始化配置文件channel
func InitConfigCh() {
	ConfigCh = make(chan RunType)
}

// InitConfig 初始化配置文件
func InitConfig() error {
	Config = new(ConfigExt)
	data, err := ioutil.ReadFile("./config/config.xml")
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, Config)
	if err != nil {
		return err
	}
	configTime = GetConfigTime()
	if configTime == -1 {
		return errors.New("获取文件信息失败")
	}
	return nil
}

// GetConfigTime 获取当前配置文件的时间
func GetConfigTime() int64 {
	f, err := os.Stat("./config/config.xml")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return f.ModTime().Unix()
}

// CheckConfig 定时检查配置文件，更新绑定按键
func CheckConfig() {
	for {
		time.Sleep(3 * time.Second)
		unix := GetConfigTime()
		if unix == -1 {
			time.Sleep(5 * time.Second)
		} else {
			if unix != configTime {
				fmt.Println("配置文件发生变化：", time.Now().Format("2006-01-02 15:04:05"))
				configTime = unix
				// 	说明文件发生变动了
				err := InitConfig()
				if err != nil {
					fmt.Println(err.Error())
					ConfigCh <- CONFIG_ERROR
					return
				}
				ConfigCh <- CONFIG_CHANGE
			}
		}

	}
}
