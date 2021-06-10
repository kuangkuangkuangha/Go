package handler

import (
	"TeamManage/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetTeamInfo 是一个根据ID获取球队信息的处理器
func GetTeamInfo(c *gin.Context) {
	var tem []int
	var teaminfo model.Team
	var player []model.Player
	var name []string
	//获取路由里的查询参数
	teamid := c.Param("id")
	//通过查询参数查数据库
	model.DB.Table("team").Where("id = ?", teamid).Find(&teaminfo)

	member := teaminfo.Member

	//  ID, err := strconv.Atoi(member)
	//  if err != nil {
	// 	 fmt.Println("error",err)
	//  }

	//  //获取球队的球员名字
	//  id := make([]int,100)
	//  if ID >= 0{
	//  i := ID%10
	//  ID = ID/10
	//  id = append(id,i)

	// }
	tem = model.GetName(member)
	model.DB.Table("players").Where("id in (?)", tem).Find(&player)
	//通过ID把球员的名字一个一个找出来

	// var temp model.Player
	for _, value := range player {
		name = append(name, value.Name)

	}
	//将含有球员姓名的切片返回给前端
	fmt.Println(name)
	c.JSON(200, gin.H{
		"message": teaminfo,
		"球员":      name,
	})
}






func GetPlayerInfo(c *gin.Context) {
	id := c.Query("player_id")
	var temp model.Player
	if err := model.DB.Table("players").Where("id = ?", id).Find(&temp).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "查询失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": temp,
		})
	}

}
