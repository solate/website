package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/solate/website/sys/mgodb"
	"gopkg.in/mgo.v2"
)

//导航
type NavigationBar struct {
	Id       int    `json:"id"`       //id
	Name     string `json:"name"`     //名字
	Url      string `json:"url"`      //引用那个页面url
	OrderNum int    `json:"ordernum"` //排序序号
	Time     int64  `json:"time"`     //创建时间
}

//根据分页获得模板
func GetAllNavigationBar(page, pageSize int) (navigationBars []NavigationBar, err error) {
	query := bson.M{

	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBNavigationBar).Find(query).Sort("ordernum").Skip(pageSize * (page - 1)).Limit(pageSize).All(&navigationBars)
	})
	return
}

//获得总数量
func GetNavigationBarCount() (allCount int, err error) {
	query := bson.M{

	}
	mgodb.Exec(func(mgosess *mgo.Session) error {
		allCount, err = mgosess.DB(DB).C(DBNavigationBar).Find(query).Count()
		return err
	})
	return
}

//添加模板
func AddNavigationBar(navigationBar *NavigationBar) (err error) {
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBNavigationBar).Insert(navigationBar)
	})
	return
}

//更新模板
func UpdateNavigationBar(navigationBar *NavigationBar) (err error) {
	query := bson.M{
		"id": navigationBar.Id,
	}
	update := bson.M{
		"$set": bson.M{
			"Name":    navigationBar.Name,
			"url":    navigationBar.Url,
			"ordernum": navigationBar.OrderNum,
		},
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBNavigationBar).Update(query, update)
	})
	return
}

//删除
func DeleteNavigationBar(id int) (err error) {
	query := bson.M{
		"id": id,
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBNavigationBar).Remove(query)
	})
	return
}

func SearchNavigationBar(name string) (navigationBars []NavigationBar, err error) {
	query := bson.M{
		"name": name,
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBNavigationBar).Find(query).Sort("ordernum").All(&navigationBars)
	})

	return
}
