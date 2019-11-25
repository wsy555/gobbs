package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func OpenDB() (*gorm.DB, error) {

	MySqlUrl := beego.AppConfig.String("MySqlUrl")
	db, err := gorm.Open("mysql", MySqlUrl)

	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}

// 关闭连接
func CloseDB() {
	if db == nil {
		return
	}
	if err := db.Close(); nil != err {
		panic(err)
	}
}

// 获取数据库链接
func GetDB() *gorm.DB {
	if db == nil {
		db, _ := OpenDB()
		return db
	}
	return db
}
