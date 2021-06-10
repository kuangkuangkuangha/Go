package handler

import (
	"TeamManage/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// func PlayerRegister(c *gin.Context) {
// 	var player model.Player
// 	if err := c.BindJSON(&player).Error; err != nil {
// 		c.JSON(404, gin.H{
// 			"message": "绑定失败",
// 		})

// 	}

// 	if err1 := model.DB.Table("players").Create(&player).Error; err1 != nil {
// 		c.JSON(404, gin.H{
// 			"message": "创建失败",
// 		})
// 	} else {

// 		c.JSON(200, gin.H{
// 			"message": "创建成功",
// 		})

// 	}
// }
//Register 是一个注册的处理器
func PlayerRegister(c *gin.Context) {

	var user model.Player
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}
	//将球员信息插入数据库
	if err1 := model.DB.Table("players").Create(&user).Error; err1 != nil {
		fmt.Println(err1)
	}
	//将球员球队关系插入数据库
	belong := model.TeamPlayer{TeamID: user.TeamBelongID, Playerid: user.ID}
	if err1 := model.DB.Table("team_players").Create(&belong).Error; err1 != nil {
		fmt.Println(err1)
	}

	c.JSON(200, gin.H{
		"username": user.ID,
	})
}

//PlayerBelongTeam 是一个登记球员球队关系的处理器
func PlayerBelongTeam(c *gin.Context) {
	var temp model.TeamPlayer
	c.BindJSON(&temp)
	if err := model.DB.Table("team_players").Create(&temp).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "插入失败",
		})
	} else {

		c.JSON(200, gin.H{
			"message": "创建成功",
		})

	}

}

func TeamResgister(c *gin.Context) {

}
