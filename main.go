package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/CYL96/MySDK/mypprof"
	"golang.org/x/net/http2"
	_ "net/http/pprof"

	"newknowlage/boom"
	"newknowlage/snake"
)

func f5() (r int) {
	defer func() {
		r++
	}()
	return 0
}
func main() {
	//fmt.Println(f5())
	mypprof.MypprofStart(true, true, 0)

	boom.Run_Mine()
	return
	//log.LogInit(log.Log_mode_Print, log.Log_Lv_Debug, log.Unix, log.Color_On)
	//RGB.Run()
	snake.SnakeInit()
	var srv http.Server
	http2.VerboseLogs = true
	srv.Addr = ":9999"
	http2.ConfigureServer(&srv, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi tester %q\n", html.EscapeString(r.URL.Path))
		ShowRequestInfoHandler(w, r)
	})
	// Listen as https ssl server
	// NOTE: WITHOUT SSL IT WONT WORK!!
	log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
}
func ShowRequestInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)
}
