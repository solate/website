package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/solate/website/sys/config"
	"github.com/solate/website/sys/mgodb"
	"github.com/solate/website/sys/logs"
	"github.com/solate/website/models"
	"gopkg.in/mgo.v2"
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

	var value =  models.Stencil{
		Image:"/images/123.jpg",
		Title:"测试1",
		Content:"content1",
		Price:"￥ 199",

	}

	dberr := mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(models.DB).C(models.DBStencil).Insert(value)
	})

	logrus.Error(dberr)



}
