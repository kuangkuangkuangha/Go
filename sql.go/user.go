package model

import "fmt"

//Register 是用户注册的函数
func Register(name string, password string, avatar string) string {

	user := User{Username: name, Password: password}
	if err := DB.Table("users").Create(&user).Error; err != nil {
		fmt.Println("registError" + err.Error())
		return " "
	}

	return user.Username
}
