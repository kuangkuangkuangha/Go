package main

import (
	"TeamManage/handler"

	"TeamManage/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db, err := gorm.Open("mysql", "root:zk2824895143@(localhost)/user?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	model.InitDB()
	r := gin.Default()

	r.POST("/user", handler.Register)
	r.PUT("/user", handler.Login)
	
	
	//路由中有查询参数（待测）
	r.POST("/team/:id",handler.GetTeamInfo)
	r.POST("/game",handler.CreateOneGame)
	r.PUT("/game",handler.ChangeTeam)
	

	r.GET("/player",handler.Getplayer)//url中有查询参数
	r.POST("/player",handler.PlayerBelongTeam)
	r.PUT("/player", handler.PlayerRegister)


	r.POST("/appoint",handler.GetGameList)
	r.GET("/appoint",handler.AppointGame)//url中有查询参数
	r.PUT("/appoint",handler.CheckAppoint)

	r.GET("/list",handler.ByTime)
	r.POST("/list",handler.ByAppoint)
	
	r.Run()

}
