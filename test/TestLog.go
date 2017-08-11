package test

import (
	"github.com/solate/website/sys/mgodb"
	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"os"
	"github.com/solate/website/sys/config"
	"github.com/solate/website/sys/logs"
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
	var value =  map[string]string{
		"test":"测试一下",
		"hhh":"nnnnnn",
	}

	dberr := mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB("test").C("user_tasks").Insert(value)
	})

	logrus.Debug(dberr)
}