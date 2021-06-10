package handler

import (
	"TeamManage/model"
	"log"
	// "fmt"

	"github.com/gin-gonic/gin"
)

//用户获取球赛列表
func GetGameList(c *gin.Context) {
	var temp  []*model.Game
	// temp = make([]*model.ListResponse,100)
	// if err:= c.BindJSON(&temp).Error;err != nil{
	// 	c.JSON(404,gin.H{
	// 		"message":"格式错误",

	// 	})
		
	// }

	if err := model.DB.Limit(2).Order("creat_at").Table("game").Find(&temp).Error; err != nil{
		log.Print(err)
	}
	
	c.JSON(200,gin.H{
		"message":temp,
	})




}
