package xsql

import (
	"strings"
)

//SQL防御，对应转义和还原
var sqlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;",
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;",
)

var sqlUnescaper = strings.NewReplacer(
	"&amp;", `&`,
	"&#39;", `'`,
	"&lt;", `<`,
	"&gt;", `>`,
	"&#34;", `"`,
)
//XSS防御
var xssEscaper = strings.NewReplacer(
	"document.","",
	"script","",
	"alert","",
	"<script>","",
	"</script>","",
	"java","",
	"exec","",
	"%2E","",
	"layer","",
	"function","fun_c",
	"link","",
	//"<","[",
	//">","]",
	//"{","[",
	//"}","]",
	//"'","~",
)

//SQL写入数据库的时候进行转义
func EncodeSQL(s string) string {
	return sqlEscaper.Replace(DecodeSQL(s))
}
//将从SQL中拿到的数据进行还原
func DecodeSQL(s string) string {
	return sqlUnescaper.Replace(s)
}
//对需要写入的数据进行过滤，防止XSS
func EncodeXSS(s string) string {
	snew := xssEscaper.Replace(strings.ToLower(s))
	//fmt.Println(len(s),len(snew))
	//if len(s)==len(snew){
	//	return s
	//}
	return snew
}