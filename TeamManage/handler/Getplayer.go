package handler

import (
	"TeamManage/model"

	"github.com/gin-gonic/gin"
)


func Getplayer(c *gin.Context) {
	id := c.Query("player_id")
	c.JSON(500, gin.H{
		"message": "查询成功，正在删除",
	})

	err := model.DB.Table("players").Where("id = ?", id).Delete(&model.Player{}).Error
	if err != nil {
		c.JSON(300, gin.H{
			"message": "查无此球员",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "删除成功",
		})
	}

}
