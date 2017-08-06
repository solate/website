package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"git.oschina.net/YPTArch/basic/sys/config"

)

var Db *DB

func InitMysql() (err error)  {
	Db,err = NewDb()
	return
}


type DB struct {
	*gorm.DB
}

func NewDb() (*DB, error) {
	dbConfig := config.Conf.DB
	//root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

