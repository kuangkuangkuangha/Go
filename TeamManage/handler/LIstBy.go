package handler

import (
	"TeamManage/model"
	"time"

	"github.com/gin-gonic/gin"
)

//用户获取球赛列表
func ByTime(c *gin.Context) {

	time := time.Now()
	var temp []model.Game

	err := model.DB.Table("game").Limit(10).Order("creat_at").Where("creat_at >?", time).Find(&temp).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "查询失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": temp,
		})
	}
}

//根据预约热度排序
func ByAppoint(c *gin.Context) {
	var temp []model.Game
	err := model.DB.Table("game").Limit(10).Order("count").Find(&temp).Error
	if err != nil {
		c.JSON(300, gin.H{
			"message": "查询失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": temp,
		})
	}

}
