package handler

import (
	"TeamManage/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//@Summary 用户预约球赛接口
//@Description
//@Tag
//@Accept
//@Produce
//@Param
//@Param
//@Success
//@Router


func AppointGame(c *gin.Context) {
	//获取请求参数中的gameid用于增加预约数
	gameid := c.Query("gameid")
	userid := c.Query("userid")

	gameID, _ := strconv.Atoi(gameid)
	userID, _ := strconv.Atoi(userid)

	usergame := model.Usergame{
		GameID: gameID,
		UserID: userID,
	}

	//将球员球赛关系插入数据库
	err := model.DB.Table("user_game").Create(&usergame).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "预约失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "预约成功",
		})
	}

	//model怎么用？？？？？
	// err := model.DB.Table("game").Where("id = ?", gameid).Update("count",count+1).Error
	err1 := model.DB.Table("game").Where("id = ?", gameid).UpdateColumn("count", gorm.Expr("count + ?", 1)).Error
	// model.DB.Model(&game)

	if err1 != nil {
		c.JSON(400, gin.H{
			"message": "预约失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "修改成功",
		})
	}

}
