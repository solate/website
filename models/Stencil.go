package models

import (
	"github.com/solate/website/sys/mgodb"
	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type Stencil struct {
	Image string `json:"image"` //模块图片
	Title string `json:"title"` //标题
	Content string `json:"content"` //内容
	Price string `json:"price"` //内容
}




func GetAllStencil(page, pageSize int) (stencils []Stencil, err error) {
	query := bson.M{

	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBStencil).Find(query).Skip(pageSize*page).Limit(pageSize).All(&stencils)
	})
	return
}

func GetStencilCount() (allCount int, err error) {
	query := bson.M{

	}
	mgodb.Exec(func(mgosess *mgo.Session) error {
		allCount, err = mgosess.DB(DB).C(DBStencil).Find(query).Count()
		return err
	})
	return
}

func StencilById()  {
	var value =  Stencil{

	}

	dberr := mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBStencil).Insert(value)
	})

	logrus.Debug(dberr)
}