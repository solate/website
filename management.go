package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/solate/website/sys/config"
	"github.com/solate/website/sys/mgodb"
	"github.com/solate/website/sys/logs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/solate/website/controllers"
)

func init() {
	//初始化相关服务器的连接
	if err := config.LoadConfig(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	//日志初始化
	if err := logs.Init(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	//初始化mgodb
	if err := mgodb.InitMgo(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

}

func main() {


	e := echo.New()

	//开发模式
	e.Debug = true

	//加载静态文件
	e.Use(middleware.Static("./views/management"))


	//路由
	e.GET("/", controllers.ManageIndex)
	//首页菜单
	e.GET("/NavigationBar/show/:page/step/:step", controllers.NavigationBarShow)
	e.POST("/NavigationBar/add", controllers.NavigationBarAdd)
	e.POST("/NavigationBar/update", controllers.NavigationBarUpate)
	e.GET("/NavigationBar/delete", controllers.NavigationBarDelete)

	//轮播图
	e.GET("/Carousel/show", controllers.NavigationBarShow)
	e.POST("/Carousel/add", controllers.NavigationBarShow)
	e.GET("/Carousel/delete", controllers.NavigationBarShow)

	//模版管理
	e.GET("/Stencil/show/:page/step/:step", controllers.StencilShow)
	e.POST("/Stencil/add", controllers.StencilAdd)
	e.DELETE("/Stencil/delete/:id", controllers.StencilDelete)
	e.POST("/Stencil/search/:title", controllers.StencilSearch)

	e.GET("/Bottom/show/:page/step/:step", controllers.StencilShow)
	e.POST("/Bottom/add", controllers.NavigationBarShow)
	e.GET("/Bottom/delete", controllers.NavigationBarShow)



	e.Start("127.0.0.1:9999")

}