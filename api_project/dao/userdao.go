package dao

import (
	"api_project/model"
    "api_project/utils"
	
)
//CheckUsernameAndPassword 是从表单中拿到用户名和密码进行校验
func CheckUsernameAndPassword(username string , password string)(*model.User , error ){
//开始到数据库里查人
//写sql语句
sqlStr :="select id , username , password, role from login where username =？and password =? "
//执行
row := utils.Db.QueryRow(sqlStr,username,password)
//
haha := &model.User{}
row.Scan(&haha.ID,&haha.Username,&haha.Password,&haha.ROle)
    return haha,nil
}