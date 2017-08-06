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

	//加载静态文件
	e.Use(middleware.Static("./views/management"))


	//路由
	e.GET("/", controllers.ManageIndex)


	e.Start("127.0.0.1:9999")



	//var value =  map[string]string{
	//	"test":"测试一下",
	//	"hhh":"nnnnnn",
	//}
	//
	//dberr := mgodb.Exec(func(mgosess *mgo.Session) error {
	//	return mgosess.DB("test").C("user_tasks").Insert(value)
	//})
	//
	//logrus.Debug(dberr)


	
}