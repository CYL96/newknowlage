package chatServer

import (
	"net"
	"fmt"
)

func ChatStart(){
	udpAddr,_ := net.ResolveUDPAddr("udp","127.0.0.1:8000")
	udpConn,_ := net.ListenUDP("udp",udpAddr)
	fmt.Println("listenAndServer:"+udpAddr.String())
	chatserver(udpConn)
}
func chatserver(conn *net.UDPConn){

	aaa ,_:= net.ResolveUDPAddr("udp","127.0.0.255:8000")
	for{
		buf := make([]byte,1024)
		_,caddr,err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf),caddr,aaa)
		n,err := conn.WriteToUDP(buf,aaa)
		fmt.Println(n,err)
	}
}
