package HTTPserver

import (
	"net"
	"fmt"
	"time"
)

func INIT(){
	var tcpAddr *net.TCPAddr
	tcpAddr,_ = net.ResolveTCPAddr("tcp","127.0.0.1:8000")
	tcpListener,_ := net.ListenTCP("tcp",tcpAddr)
	fmt.Println("listenAndServer:"+tcpAddr.String())
	defer tcpListener.Close()
	for{
		tcpConn,err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("connected failed.")
			continue
		}
		fmt.Println("A client connected:"+tcpConn.RemoteAddr().String())
		go tcpdeal(tcpConn)
	}
}
func tcpdeal(conn *net.TCPConn){
	ipstr := conn.RemoteAddr().String()
	defer func(){
		fmt.Println("disconnected :"+ipstr)
		conn.Close()
	}()
	buf := make([]byte,1024)
	for{
		_,err := conn.Read(buf)

		if err != nil{
			fmt.Println(err.Error())
			fmt.Println(ipstr,"--closed")
			return
		}
		fmt.Println("VVVVVVVVVVV",ipstr,"VVVVVVVVVVV")
		fmt.Println(string(buf))
		fmt.Println("AAAAAAAAAAA",ipstr,"AAAAAAAAAAA")
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}

}