package handler

import (
	"TeamManage/model"

	"github.com/gin-gonic/gin"
)

//AppointGame 是一个预约球赛的处理器函数
func AppointGame(c *gin.Context) {
	var temp model.Appoint
	if err := c.BindJSON(&temp).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}
	c.JSON(200, gin.H{
		"mesage": "插入成功",
	})

}
