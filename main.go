/**
 * Auth :   liubo
 * Date :   2018/10/23 9:47
 * Comment: 
 */

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

var addr = flag.String("addr", ":8392", "addr=ip:port")

func main() {
	flag.Parse()
	fmt.Println("启动服务器，地址是：", *addr)

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/myip", func(c* gin.Context) {

		ip, port, _ := net.SplitHostPort(c.Request.RemoteAddr)

		c.String(http.StatusOK, "ip+port:" + ip + ":" + port)
	})

	r.Run(*addr)

	fmt.Println("关闭服务器")
}
