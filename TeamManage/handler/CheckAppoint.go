package handler

import (
	"TeamManage/model"
	"fmt"
	"log"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"
)

//用户查看自己预约球赛的接口
func CheckAppoint(c *gin.Context) {

	id, _ := strconv.Atoi(c.Query("user_id"))
	// id := c.Query("user_id")
	// temp := make([]int,100) //球赛ID切片
	// name := make([]string, 10)
	// var temp struct{
	// 	Id int `gorm:"game_id"`
	// }
	var name []model.Usergame
	var temp []int
	var game2 []model.Game
	var name2 []string
	
	// var gamename []string
	// var name []int
	
	
	err := model.DB.Table("user_game").Where("user_id=?",id).Find(&name).Error
	for _,tempid := range name{
		temp = append(temp, (tempid.GameID))
	}
	fmt.Println(temp)

	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{
			"message": "查找失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": name,
		})
	}

// 根据用户Id查他的球赛名字
//	for i := 0;i <len(name); i++{

		model.DB.Table("game").Where("id in (?)", temp).Find(&game2)
		

	//	name2 = append(name2, game2.Name)
		
//	}


	// for _, value := range name {
	// 	fmt.Println(value.GameID)
	// 	model.DB.Table("game").Where("id = ?", value.GameID).Find(&game2)
	// 	game = append(game, game2)
	// 	fmt.Println(game)
	// 	// c.JSON(200, gin.H{

	// 	// })
	// }
		fmt.Println(game2)
		for _ , value := range game2{
			name2 = append(name2,value.Name)
		}

	// for _, value := range game{
	// 	gamename = append(gamename, value.Name)
	// }

	c.JSON(200, gin.H{
		"message":name2,
	})

}
