package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

var count = 0

// SetEngineHandle 设置
func SetEngineHandle(engine *gin.Engine) {
	engine.GET("/getAddr", func(c *gin.Context) {
		count += 1
		log.Printf("call:%d\n", count)
		resp := fmt.Sprintf("IP: %s\nTotalReq: %d", GetLocalIPDotFormat(), count)
		c.String(http.StatusOK, resp)
	})
}

func GetLocalIPDotFormat() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("get local ip failed, error: %s\n", err)
		return ""
	}

	for _, address := range addrs {
		if ipNet, ok := address.(*net.IPNet); ok {
			if isIPUseful(ipNet.IP) {
				return ipNet.IP.String()
			}
		}
	}

	log.Println("no valid ipv4")
	return ""
}

func isIPUseful(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	ipv4 := ip.To4()
	if ipv4 == nil {
		return false
	}
	return true
}
