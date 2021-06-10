package main

import (
	"bookstore/model"
	"database/sql"
	"fmt"
	"net/http"
)

// Db shi
var Db *sql.DB

func connectmysql() (err error) {
	Db, err = sql.Open("mysql", "root:zk2824895143@tcp(localhost:3306)/user")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}

func main() {
	err := connectmysql()
	if err != nil {
		fmt.Printf("init DB failed,err:5v%\n", err)
	}
	fmt.Println("数据库连接成功")

	http.HandleFunc("/login", model.Login)
	http.ListenAndServe("8080:", nil)
}
