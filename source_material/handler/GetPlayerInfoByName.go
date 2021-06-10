package handler

import (
	"TeamManage/model"
	"log"

	"github.com/gin-gonic/gin"
)

//先声明一个结构体
//然后Bindjson从表单中找信息并初始化结构体
//如果绑定失败则判断错误
//通过写好的函数得到要返回给前端的字段
//通过c.json将数据返回给前端





//GetPlayerInfoByName 是一个根据球员姓名获取球员信息的处理器
func GetPlayerInfoByName(c *gin.Context) {
	var info model.Player
	if err := c.BindJSON(&info); err != nil {
		log.Println(err)
		c.JSON(401,gin.H{
			"message":"输入格式有误"})
	} 
	result, err := model.GetPlayerInfoByName(info.Name)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "没有这个球员的信息"})
		return
	}
	c.JSON(200, result)

}
