package core

import (
	"encoding/xml"
	"io/ioutil"
)

var Config ConfigT

type ConfigT struct {
	XMLName xml.Name `xml:"config"`
	Color   ColorT   `xml:"color"` //日志配置;
	Info    InfoT    `xml:"info"`
}
type ColorT struct {
	XMLName xml.Name   `xml:"color"`
	Title   ColorRGBAT `xml:"title"`
	Content ColorRGBAT `xml:"content"`
	Back    ColorRGBAT `xml:"back"`
}
type ColorRGBAT struct {
	R uint8 `xml:"r"`
	G uint8 `xml:"g"`
	B uint8 `xml:"b"`
	A uint8 `xml:"a"`
}

type InfoT struct {
	XMLName xml.Name `xml:"info"`
	Desc    []DescT  `xml:"desc"`
}
type DescT struct {
	XMLName  xml.Name `xml:"desc"`
	Content  string   `xml:"content"`
	Size     int      `xml:"size"`
	Kerning  int      `xml:"kerning"`
	Line     int      `xml:"line"`
	BackSize int      `xml:"back_size"`
}

func InitConfig() error {
	data, err := ioutil.ReadFile("./config.xml")
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, &Config)
}
