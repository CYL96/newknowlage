package xsql

import "strings"

var escaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;",
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;",
)

var unescaper = strings.NewReplacer(
	"&amp;", `&`,
	"&#39;", `'`,
	"&lt;", `<`,
	"&gt;", `>`,
	"&#34;", `"`,
)

func EncodeSQLandXSS(s string) string {
	return escaper.Replace(s)
}

func DecodeSQLandXSS(s string) string {
	return unescaper.Replace(s)
}
