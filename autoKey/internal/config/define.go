package config

import "encoding/xml"

type ConfigExt struct {
	XMLNAme xml.Name    `xml:"config"`
	HotKey  []HotKeyExt `xml:"hot_key"`
	Run     RunExt      `xml:"run"`
}

type RunExt struct {
	XMLNAme   xml.Name `xml:"run"`
	AfterTime int      `xml:"after_time"`
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
	CONFIG_ERROR  RunType = -1
	CONFIG_CHANGE RunType = 1
)
