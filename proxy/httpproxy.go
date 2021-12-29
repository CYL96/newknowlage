package proxy

import (
	"fmt"
	"net"

	"bytes"
	"io"
	"net/url"
	"strings"
)

func LListener() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		go handleDeal(conn)
	}
}

func handleDeal(conn net.Conn) {
	request(conn)
}

func request(conn net.Conn) (err error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &host)
	//fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &host)
	fmt.Println(string(buf), "\n", buf, "\n", byte('\n'))
	fmt.Println(method, "-", host)
	fmt.Println(string(buf))
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
	server, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(buf[:n])
	}
	//进行转发
	go io.Copy(server, conn)
	io.Copy(conn, server)
	return
}

//http://127.0.0.1:8081
