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

func GetConfig() *ConfigExt {
	return Config
}
func InitConfigCh() {
	ConfigCh = make(chan RunType)
}
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

func GetConfigTime() int64 {
	f, err := os.Stat("./config/config.xml")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return f.ModTime().Unix()
}

func CheckConfig() {
	for {
		time.Sleep(3 * time.Second)
		unix := GetConfigTime()
		if unix == -1 {
			time.Sleep(5 * time.Second)
		} else {
			if unix != configTime {
				fmt.Println("配置文件发生变化：")
				configTime = unix
				// 	说明文件发生变动了
				err := InitConfig()
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(2)
				}
				ConfigCh <- CONFIG_CHANGE
			}
		}

	}
}
