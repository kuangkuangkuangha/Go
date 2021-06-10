package handler

import (
	"TeamManage/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

//Register 是一个注册的处理器
func Register(c *gin.Context) {

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}

	if err1 := model.DB.Table("users").Create(&user).Error; err1 != nil {
		fmt.Println(err1)
	}

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

//Login 是一个登陆的函数
func Login(c *gin.Context) {

	var user model.User
	var user2 model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}

	model.DB.Table("users").Where("username = ? AND password = ?", user.Username, user.Password).Find(&user2)
	if user2.ID > 0 {
		c.JSON(200, gin.H{
			"message":  "登入成功",
			"username":user2.Username,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "登入失败",
		})
	}


	token := model.CreateToken(user.ROle)
	fmt.Println(token)

	c.JSON(200, gin.H{
		"message": token,
	})



}
