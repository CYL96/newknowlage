package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type ConfigExt struct {
	Suf string `json:"suf" type:"varchar(255)" default:"''" comment:""`
	Pre string `json:"pre" type:"varchar(255)" default:"''" comment:""`
}

func main() {
	var config ConfigExt
	fmt.Println("请输入匹配的后缀名(不输入全匹配)：")
	fmt.Scanln(&config.Suf)
	fmt.Println("请输入需要修改的前缀名：")
	fmt.Scanln(&config.Pre)
	fList, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range fList {
		if entry.IsDir() {
			continue
		}
		suf := path.Ext(entry.Name())
		if suf == ".exe" {
			continue
		}
		if (config.Suf == "" || config.Suf == suf) && strings.Index(entry.Name(), config.Pre) != 0 {
			fmt.Println("修改文件：", entry.Name(), " ", config.Pre+entry.Name())
			os.Rename(entry.Name(), config.Pre+entry.Name())
		} else {
			fmt.Println("跳过文件：", entry.Name())
		}
	}
	fmt.Println("输入回车结束")
	fmt.Scanln(&config.Pre)
}
