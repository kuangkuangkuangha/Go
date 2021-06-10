package handler

import (
	"TeamManage/model"
	"log"
	// "strconv"

	"github.com/gin-gonic/gin"
)

//用户查看自己预约球赛的接口
func CheckAppoint(c *gin.Context){

	// id ,_ := strconv.Atoi(c.Param("user_id"))
	id := c.Query("user_id")
	// temp := make([]int,100) //球赛ID切片
	// name := make([]string, 10)
	var temp struct{
		Id int `gorm:"game_id"`
	}
	// var name []string
	//会报错
	err := model.DB.Table("user_game").Where("id = ?", id).Select("game_id").Find(&temp)

	if err != nil {
		log.Print(err)
		c.JSON(400, gin.H{
			"message": "查找失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": temp,

		})
	}


	// for _, value := range temp {

	// 	model.DB.Table("game").Where("id = ?", value).Select("name").Find(&name)


	// }

	// c.JSON(200, gin.H{
	// 	"message":name,
	// })



}