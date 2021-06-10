package model

import "log"


//通过团队 id 获取单个团队的信息，如 logo 、名字、信息等
func GetTeamInfoById(id int)Team{
	var temp Team
//如果没查到就返回错误，并打印日志
	if err := DB.Table("team").Where("id = ?",id).Find(&temp).Error;err != nil{
		log.Println(err)
	}
//如果查到了，就保存在temp中
	return temp


}