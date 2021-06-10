package handler

import (
	"TeamManage/model"

	"github.com/gin-gonic/gin"
)

//CreatAagame创建一场球赛
func CreatAagame(c gin.Context) {
	var temp model.ListResponse
	c.BindJSON(&temp)
	err := model.DB.Create(&temp).Error
	if err != nil {
		c.JSON(404, gin.H{
		"message": "创建失败",
		})
	}
	c.JSON(200,gin.H{
		"message": "success",
	})

}
