package mgodb

import (
	"gopkg.in/mgo.v2"
	"time"
	"github.com/Sirupsen/logrus"
	"github.com/solate/website/sys/config"
)




//global
var globalMgoSession *mgo.Session


var MongoDBHosts = []string{"127.0.0.1:27017"}

func init()  {
	//如果有配置就用配置文件中的
	if len(config.Conf.MongoDBHosts) > 0  {
		MongoDBHosts = config.Conf.MongoDBHosts
	}
}

func InitMgo() (err error)  {
	dialInfo := &mgo.DialInfo{
		Addrs: MongoDBHosts,
		Timeout: time.Second * 3,
		PoolLimit: 128,
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