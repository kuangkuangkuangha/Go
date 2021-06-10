package handler

import (
	"TeamManage/model"

	"github.com/gin-gonic/gin"
)

//创建一个球赛
func CreateOneGame(c *gin.Context) {

	//首先用 c.Request.Header.Get 从前端那里得到token
	//然后用写好的函数 .VerifyToken(token) 解析得到一个实例化结构体



	token := c.Request.Header.Get("token")
	role, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	
	if role !=0{
		c.JSON(404,gin.H{
			"message": "用户无权限",
		})
	}else{







	var temp model.Game
	c.BindJSON(&temp)

	if err := model.DB.Table("game").Create(&temp).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "创建失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "创建成功",
		})
	}
}


}

