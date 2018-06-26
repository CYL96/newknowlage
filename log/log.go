package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
	"time"
)

var config map[string]interface{} = make(map[string]interface{})

func Init(autoclean bool) {

	readcongfig()
	createlogfile()
	if autoclean {
		go timer()
	}
}

//定时器
func timer() {
	tk := time.NewTicker(time.Hour * 8)
	go cleanfile()
	for {
		select {
		case <-tk.C:
			go cleanfile()
		}
	}

}

func WriteErr(err error) {
	tm := time.Now()
	fmt.Println("/nVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVV/n")
	fmt.Println("/n>>>>>>>>>>>", tm)
	fmt.Println(err)
	fmt.Println(string(debug.Stack()))
	fmt.Println("/nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/n")
}

func cleanfile() {
	now := time.Now()
	expire := 7.0
	if config["expire"] != "" {
		expire = config["expire"].(float64)
	}
	path := "./log"
	if config["path"] != "" {
		path = config["path"].(string)
	}
	//todo 可以封装成一个函数，作为读取摸个目录下的文件
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range dir {
		if v.IsDir() {
			continue
		}
		if v.ModTime().Add(time.Hour * 24 * time.Duration(expire)).Before(now) {
			fmt.Println(v.Name(), "已移除\r\n")
			os.Remove(path + "/" + v.Name())
		}
	}

}

//读取配置文件
func readcongfig() {
	fi, err := os.Open("./log_config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	data, err := ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("config文件错误")
		return
	}
	return
}

//创建log文件
func createlogfile() {
	t1 := time.Now().Format("2006_01_02_15_04_05")
	fmt.Println(t1)
	path := "./log"
	if config["path"] != "" {
		path = config["path"].(string)
	}
	pathcreate(path)

	path += "/" + t1 + ".dat"

	logfile, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	os.Stdout = logfile
	os.Stderr = logfile
	return
}

//自动创建文件夹
func pathcreate(path string) error {
	var err error
	_, err = os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.MkdirAll(path, 0711)
	}
	return err
}
