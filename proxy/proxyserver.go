package proxy

import (
	"net"
	"fmt"
	"bytes"
	"net/url"
	"strings"
	"io"
)

func StartProxy(){
	l,err:= net.Listen("tcp",":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l.Addr().String())
	for {
		client, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleClientRequest(client)
	}
}
func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	fmt.Println(method,"=",host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
		return
	}
	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"

		} else {
			address = hostPortURL.Host
		}

	}
	fmt.Println(address)
	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	}
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}