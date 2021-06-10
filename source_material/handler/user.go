package handler

import (
	"TeamManage/middlewale"
	"TeamManage/model"
	"fmt"

	"github.com/gin-gonic/gin"
)


//先声明一个结构体
//然后Bindjson从表单中找信息并初始化结构体
//如果绑定失败则判断错误
//





//Register 是一个注册的处理器
func Register(c *gin.Context) {

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}

	username := model.Register(user.Name, user.Password, user.Avatar)

	fmt.Println(username)
	c.JSON(200, gin.H{
		"username": username,
	})
}

//Login 是用于登入的函数
func Login(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入格式错误"})
		return
	}
	fmt.Println(user.Name, user.Password)
	// 验证用户是否存在
	userID := model.Login(user.Name, user.Password)
	if userID > 0{
		c.JSON(400, gin.H{
			"message": "登入成功",

		})

		//并生成Token
		token := middlewale.SetToken(user.Name)
		//将Token返回
		c.JSON(200,gin.H{
			"Token" :token,
		})
	}
}










//PlayerRegister 是一个球员登记的处理器
func PlayerRegister(c *gin.Context) {
	var player model.Player
	if err := c.BindJSON(&player); err != nil {
		c.JSON(400,gin.H{
			"message": "输入有误，格式错误"})
			return
	}

	//这一步是不是多余的？？？？？？
	playername := model.Registerplayer(player.Name, player.Picture, player.TeamBelong)

fmt.Println(playername)

	c.JSON(200, gin.H{
		"playername": playername,
	})
}

//Registerteam 是用于球队登记的处理器
 func Registerteam(c *gin.Context){
	 var team model.Team
	 if err := c.BindJSON(&team); err != nil {
		 c.JSON(400,gin.H{
			 "message": "输入有误，格式错误"})
			 return 
	 }
	 teamname := team.Name
	 fmt.Println(teamname)

	 c.JSON(200,gin.H{
		 "teamname": teamname })

 }
