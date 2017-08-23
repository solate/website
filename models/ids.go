package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/solate/website/sys/mgodb"
	"errors"
)

//自增长主键
type Ids struct {
	Name string `json:"name"` //名字
	Id   int `json:"id"`      //id
}

//获得唯一id
func GetId(name string) (id int) {
	var ids []Ids
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		Upsert:    true,
		ReturnNew: false,
	}
	err := mgodb.Exec(func(mgosess *mgo.Session) error {
		_, err := mgosess.DB(DB).C(DBIds).Find(bson.M{"name": name}).Apply(change, &ids)
		return err

	})
	CheckError(err)

	if len(ids) == 1 {
		id = ids[0].Id
	} else {
		CheckError(errors.New(name + " get ids err:" + err.Error()))
	}

	return
}
