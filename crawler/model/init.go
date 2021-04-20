package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB 是一个连接池
var DB *gorm.DB
var err error

//InitDB 是一个初始化的函数
func InitDB() *gorm.DB {
	DB, err = gorm.Open("mysql", "root:zk2824895143@(localhost)/user?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	//defer DB.Close()
	return DB
}
