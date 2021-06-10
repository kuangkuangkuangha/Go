package handler

import (
	"TeamManage/model"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

//球员换队伍
func ChangeTeam(c *gin.Context) {
	id ,_:= strconv.Atoi(c.Query("player_id")) 
	newteamid ,_:= strconv.Atoi(c.Query("team_id"))

	var temp model.TeamPlayer
	
	model.DB.Table("team_players").Where("playerid = ?", id).Find(&temp)

	fmt.Println(temp)
	//多余一步
	// err := model.DB.Model(&temp).Where("playerid = ?", id).Update("team_id", newteamid)
	//会报错
	temp.TeamID=newteamid

	err:=model.DB.Table("team_players").Save(&temp).Error


	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{
			"message": "换队失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "换队成功",
		})
	}
}
