package handler

import (
	"TeamManage/model"

	"github.com/gin-gonic/gin"
)

func ChangeRole(c *gin.Context){
//验证权限
	token := c.Request.Header.Get("token")
	role, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	if role != 2{
		return
	}


//最高权限者给他们权限
	var temp model.User

	user_id := c.Query("user_id")
	role2 := c.Query("role")

	if err := model.DB.Table("Users").Where("id = ?", user_id).Find(&temp).Error;err != nil{
		c.JSON(400, gin.H{
			"message": "输入有wu"})
		return
	}

	err1 := model.DB.Table("Users").Update("role",role2).Error; if err1 != nil{
		c.JSON(400, gin.H{
			"message": "授权失败"})
		return
	}


}

