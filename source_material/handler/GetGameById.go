package handler

import (
	"TeamManage/model"
	"log"

	"github.com/gin-gonic/gin"
)

//GetGameByID 是一个通过ID获取球员信息的处理器
func GetGameByID(c *gin.Context){
	var temp model.Team
	if err := c.BindJSON(&temp);err != nil {
		log.Println(err)
		c.JSON(404,gin.H{
			"message":"输入格式有误",
		})
	}
	res := model.GetGameById(temp.ID)

	c.JSON(200,res)
}
