package translate

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
 * Created by Vim
 * User: lock
 * Date: 2018/10/25
 * Time: 14:11
 * github ：github.com/LockGit/Go/blob/master/go-dict/dict.go
 */

type ParseXmlData struct {
	XMLName     xml.Name     `xml:"yodaodict"`
	RawInput    string       `xml:"return-phrase"`
	CustomTrans CustomNode   `xml:"custom-translation"`
	WebTrans    WebTransList `xml:"yodao-web-dict"`
}

type CustomNode struct {
	Type        string        `xml:"type"`
	Translation []Translation `xml:"translation"`
}

type WebTransList struct {
	TransNode []WebTransNode `xml:"web-translation"`
}

type WebTransNode struct {
	Key   string      `xml:"key"`
	Trans []TransNode `xml:"trans"`
}

type TransNode struct {
	Value string `xml:"value,CDATA"`
}

type Translation struct {
	Content string `xml:"content,CDATA"`
}

func HttpGet(url string, ch chan []byte) chan []byte {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("request error:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io read error:", err)
	}
	ch <- body
	return ch
}

func Translate(worlds ...string) string {

	word := strings.Join(worlds, " ")
	requestUrl := fmt.Sprintf("http://dict.youdao.com//fsearch?client=deskdict&keyfrom=chrome.extension&q=%s&pos=-1&doctype=xml&xmlVersion=3.2&dogVersion=1.0&vendor=unknown&appVer=3.1.17.4208&le=eng", url.QueryEscape(word))
	ch := make(chan []byte)
	go HttpGet(requestUrl, ch)

	xmlObject := ParseXmlData{}
	xml.Unmarshal(<-ch, &xmlObject)
	//fmt.Println("您的输入:", strings.TrimSpace(xmlObject.RawInput))
	var buf bytes.Buffer
	buf.WriteString("*******英汉翻译:*******\r\n")
	for _, v := range xmlObject.CustomTrans.Translation {
		buf.WriteString(strings.TrimSpace(v.Content))
		buf.WriteString("\r\n")
	}
	buf.WriteString("*******英汉翻译:*******\r\n")
	for _, v := range xmlObject.WebTrans.TransNode {
		key := strings.TrimSpace(v.Key)
		for _, vv := range v.Trans {
			value := strings.TrimSpace(vv.Value)
			buf.WriteString(key)
			buf.WriteString(":")
			buf.WriteString(value)
			buf.WriteString("\r\n")

			break
		}
	}
	return buf.String()
}
