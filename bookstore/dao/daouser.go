package dao

import (
	"api_project/model"
	"api_project/utils"
)

//CheckUsernameAndPassword 函数
func CheckUsernameAndPassword(username string, password string) *model.User {
	sqlStr := "select id , username , password, role from login where username = ? and password =?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	a := &model.User{}
	row.Scan(&a.ID, &a.Password, &a.Username, &a.ROle)
	return a
}

//CheckUser 函数
func CheckUser(username string) (a *model.User) {
	sqlStr := "select id , username , password, role from login where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	a = &model.User{}
	row.Scan(&a.ID, &a.Password, &a.ROle, &a.Username)
	return a
}
