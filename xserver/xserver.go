package xserver

import (
	"github.com/gin-gonic/gin"
	"net"
	"os"
)

type XServer struct {
	address string
	server *gin.Engine
	listener net.Listener
	signalChan chan os.Signal
	Routes map[string]gin.HandlerFunc

}


func (this *XServer)RunServer2(){
	gin.SetMode(gin.DebugMode)
	this.server.GET("/test")
}