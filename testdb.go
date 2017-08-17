package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/solate/website/models"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

//global
var globalMgoSession *mgo.Session


var MongoDBHosts = []string{"127.0.0.1:27017"}

func InitMgo() (err error)  {
	dialInfo := &mgo.DialInfo{
		Addrs: MongoDBHosts,
		Timeout: time.Second * 3,
		PoolLimit: 128,
		Username:  "root",
		Password:  "iamapassword",
	}
	globalMgoSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Errorf("CreateSession: %s\n", err)
		return
	}

	globalMgoSession.SetMode(mgo.Monotonic, true)

	return

}
func Exec(f func(*mgo.Session) error) error {
	sessionCopy := globalMgoSession.Copy()
	defer sessionCopy.Close()
	return f(sessionCopy)
}

func main() {

	//初始化
	err := InitMgo()
	if err != nil {
		fmt.Println(err, "-------1---------")
	}

	var value =  models.Stencil{
		Image:"/images/123.jpg",
		Title:"测试1",
		Content:"content1",
		Price:"￥ 199",

	}

	err = Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(models.DB).C(models.DBStencil).Insert(value)
	})

	logrus.Error(err)

	var stencils []models.Stencil
	err = Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(models.DB).C(models.DBStencil).Find(bson.M{}).Skip(0).Limit(10).All(&stencils)
	})

	fmt.Println(stencils, err)



}
