package model

import "fmt"

//AppointGame 是一个预约比赛的函数
func AppointGame(userid int, gameid int) error{
	temp := Appoint{UserID: userid, GameID: gameid}
	if err := DB.Table("Appoint").Create(&temp).Error; err != nil {
		fmt.Println("registError" + err.Error())
		return err
	}
	return nil


}
