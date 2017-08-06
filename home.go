package main

import (
	"github.com/labstack/echo"
	"github.com/solate/website/controllers"
	"github.com/labstack/echo/middleware"
	"strings"
	"net/http"
	"github.com/labstack/gommon/log"
	"github.com/solate/website/sys/logs"
)

func init() {
	logs.Init()
}

func main() {
	e := echo.New()

	//加载静态文件
	e.Use(middleware.Static("./views/home"))


	//路由
	e.GET("/", controllers.Home)


	e.Start("127.0.0.1:8888")
}

//打印ip中间件
func PrintIp() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Info(IP(c.Request())) //打印ip
			return next(c)
		}
	}
}

// IP returns request client ip.
// if in proxy, return first proxy id.
// if error, return 127.0.0.1.
func IP(request *http.Request) string {
	ips := Proxy(request)
	if len(ips) > 0 && ips[0] != "" {
		rip := strings.Split(ips[0], ":")
		return rip[0]
	}
	ip := strings.Split(request.RemoteAddr, ":")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}

// Proxy returns proxy client ips slice.
func Proxy(request *http.Request) []string {
	if ips := request.Header.Get("X-Forwarded-For"); ips != "" {
		return strings.Split(ips, ",")
	}
	return []string{}
}