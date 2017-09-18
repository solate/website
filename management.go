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
	//导航
	e.GET("/NavigationBar/show/:page/step/:step", controllers.NavigationBarShow)
	e.POST("/NavigationBar/add", controllers.NavigationBarAdd)
	e.DELETE("/NavigationBar/delete/:id", controllers.NavigationBarDelete)
	e.POST("/NavigationBar/search/:name", controllers.NavigationBarSearch)

	//轮播图
	e.GET("/Carousel/show/:page/step/:step", controllers.CarouselShow)
	e.POST("/Carousel/add", controllers.CarouselAdd)
	e.DELETE("/Carousel/delete/:id", controllers.CarouselDelete)
	e.POST("/Carousel/search", controllers.CarouselSearch)

	//模版管理
	e.GET("/Stencil/show/:page/step/:step", controllers.StencilShow)
	e.POST("/Stencil/add", controllers.StencilAdd)
	e.DELETE("/Stencil/delete/:id", controllers.StencilDelete)
	e.POST("/Stencil/search/:title", controllers.StencilSearch)

	//底部
	//e.GET("/Stencil/show/:page/step/:step", controllers.StencilShow)
	//e.POST("/Stencil/add", controllers.StencilAdd)
	//e.DELETE("/Stencil/delete/:id", controllers.StencilDelete)
	//e.POST("/Stencil/search/:title", controllers.StencilSearch)



	// ------------ 用户部分 --------------
	e.GET("/login", controllers.Login)
	e.GET("/register", controllers.Register)



	e.Start("127.0.0.1:9999")

}