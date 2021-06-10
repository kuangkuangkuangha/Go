package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB 是一个连接池
var DB *gorm.DB
var err error

//InitDB 是一个初始化的函数
func InitDB() *gorm.DB {
	db, _ := gorm.Open("mysql", "haha:zk2824895143@(localhost)/?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	return DB
}
