package model

import "fmt"

//Register 是用户注册的函数
func Register(name string, password string, avatar string) string {

	user := User{Name: name, Password: password}
	if err := DB.Table("user").Create(&user).Error; err != nil {
		fmt.Println("registError" + err.Error())
		return " "
	}

	return user.Name
}


//Login 是用户登入的函数，从数据库中查取记录
func Login(name string, password string) int {
	var user User
	err := DB.Where("Name = ? AND Password = ?", name, password).Find(&user).Error
	if err != nil {
		return 0
	}
	return int(user.ID)
}

//