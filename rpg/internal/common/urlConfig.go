package common

import (
	"encoding/xml"
	"net/url"
	"os"
)

type (
	UrlConfigExt struct {
		XMLName xml.Name `xml:"url_config"`
		UrlList []UrlExt `xml:"url_ext"`
	}
	UrlExt struct {
		XMLName xml.Name `xml:"url_ext"`
		Name    string   `xml:"name"`
		Url     string   `xml:"url"`
		UrlS    *url.URL
	}
)

var UrlConfig *UrlConfigExt

func InitUrlConfig() (err error) {
	var (
		data []byte
	)

	data, err = os.ReadFile("./config/url.xml")
	if err != nil {
		return
	}
	UrlConfig = new(UrlConfigExt)
	err = xml.Unmarshal(data, UrlConfig)
	if err != nil {
		return
	}
	for i, ext := range UrlConfig.UrlList {
		UrlConfig.UrlList[i].UrlS, _ = url.Parse(ext.Url)
	}
	return
}

// GetUrlConfig 获取url配置
func GetUrlConfig() UrlConfigExt {
	if UrlConfig == nil {
		return UrlConfigExt{}
	}
	return *UrlConfig
}
