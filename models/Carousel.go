package models

import (
	"github.com/solate/website/sys/mgodb"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

//轮播图
type Carousel struct {
	Id        int    `json:"id"`        //id
	Image     string `json:"image"`     //模版图片
	Title     string `json:"title"`     //标题
	Content   string `json:"content"`   //内容
	OrderNum  int    `json:"ordernum"`  //排序序号
	IsDisplay bool   `json:"isdisplay"` //是否显示
	Time      int64  `json:"time"`      //创建时间
}

//根据分页获得模板
func GetAllCarousel(page, pageSize int) (carousels []Carousel, err error) {
	query := bson.M{

	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBCarousel).Find(query).Sort("ordernum").Skip(pageSize * (page - 1)).Limit(pageSize).All(&carousels)
	})
	return
}

//获得总数量
func GetCarouselCount() (allCount int, err error) {
	query := bson.M{

	}
	mgodb.Exec(func(mgosess *mgo.Session) error {
		allCount, err = mgosess.DB(DB).C(DBCarousel).Find(query).Count()
		return err
	})
	return
}

//添加模板
func AddCarousel(carousel *Carousel) (err error) {
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBCarousel).Insert(carousel)
	})
	return
}

//更新模板
func UpdateCarousel(carousel *Carousel) (err error) {
	query := bson.M{
		"id": carousel.Id,
	}
	update := bson.M{
		"$set": bson.M{
			"image":     carousel.Image,
			"title":     carousel.Title,
			"content":   carousel.Content,
			"ordernum":  carousel.OrderNum,
			"isdisplay": carousel.IsDisplay,
		},
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBCarousel).Update(query, update)
	})
	return
}

//删除
func DeleteCarousel(id int) (err error) {
	query := bson.M{
		"id": id,
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBCarousel).Remove(query)
	})
	return
}

func SearchCarousel(image string) (carousels []Carousel, err error) {
	query := bson.M{
		"image": image,
	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBCarousel).Find(query).Sort("ordernum").All(&carousels)
	})

	return
}

//根据分页获得模板
func GetCarouselList() (carousels []Carousel, err error) {
	query := bson.M{

	}
	err = mgodb.Exec(func(mgosess *mgo.Session) error {
		return mgosess.DB(DB).C(DBCarousel).Find(query).Sort("ordernum").All(&carousels)
	})
	return
}
