package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/solate/website/sys/mgodb"
)

type Stencil struct {
	Id       int `json:"id"`          //id
	Image    string `json:"image"`    //模版图片
	Title    string `json:"title"`    //标题
	Content  string `json:"content"`  //内容
	Price    string `json:"price"`    //价格
	OrderNum int `json:"ordernum"` //排序序号
	Time     int64  `json:"time"`     //创建时间
}

//根据分页获得模板
func GetAllStencil(page, pageSize int) (stencils []Stencil, err error) {
	query := bson.M{

	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBStencil).Find(query).Sort("ordernum").Skip(pageSize * (page - 1)).Limit(pageSize).All(&stencils)
	})
	return
}

//获得总数量
func GetStencilCount() (allCount int, err error) {
	query := bson.M{

	}
	mgodb.Exec(func(mgosess *mgo.Session) error {
		allCount, err = mgosess.DB(DB).C(DBStencil).Find(query).Count()
		return err
	})
	return
}

//添加模板
func AddStencil(stencil *Stencil) (err error) {
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBStencil).Insert(stencil)
	})
	return
}

//更新模板
func UpdateStencil(stencil *Stencil) (err error) {
	query := bson.M{
		"id": stencil.Id,
	}
	update := bson.M{
		"$set": bson.M{
			"image":    stencil.Image,
			"title":    stencil.Title,
			"content":  stencil.Content,
			"price":    stencil.Price,
			"ordernum": stencil.OrderNum,
		},
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBStencil).Update(query, update)
	})
	return
}
//删除
func DeleteStencil(id int) (err error) {
	query := bson.M{
		"id": id,
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBStencil).Remove(query)
	})
	return
}