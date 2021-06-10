package utils

import (
	_"github.com/go-sql-driver/mysql"
		"database/sql"
)

// Db 是一个连接池
var Db *sql.DB
//Connectmysql 是连接数据库
func Connectmysql()(err error){
	Db, err = sql.Open("mysql", "root:zk2824895143@tcp(localhost:3306)/user")
	if err != nil{
	return
	}
	err =Db.Ping()
	if err != nil{
	  return
	}
	return
}	