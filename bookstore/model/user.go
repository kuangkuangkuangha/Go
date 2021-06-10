package model

import (
	"bookstore/dao"
	"fmt"
	"net/http"
)

//User 是一个结构体
// type User struct {
// 	Ida      string
// 	Username string
// 	Password string
// 	Role     string
// }

// Login 登陆的处理器终于写好了！！！
func Login(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	ret := dao.CheckUsernameAndPassword(username, password)
	if ret.ID != 0 {
		fmt.Fprintln(w, "登陆成功")
	} else {
		fmt.Fprintln(w, "登陆失败")
	}
}
