package handler

import (
	"TeamManage/model"

	"github.com/gin-gonic/gin"
)


func GetAllGames(c *gin.Context){
	var listResponse  model.ListResponse
	c.BindJSON(&listResponse)
}