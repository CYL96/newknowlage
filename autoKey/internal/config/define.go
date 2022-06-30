package config

import "encoding/xml"

type ConfigExt struct {
	XMLNAme xml.Name    `xml:"config"`
	HotKey  []HotKeyExt `xml:"hot_key"`
}

type HotKeyExt struct {
	XMLNAme xml.Name `xml:"hot_key"`
	Name    string   `xml:"name"`
	Key     string   `xml:"key"`
	Account string   `xml:"account"`
	Pwd     string   `xml:"pwd"`
}

type RunType int

const (
	CONFIG_CHANGE RunType = 1
)
