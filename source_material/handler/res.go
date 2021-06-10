package handler


func ListResponse()(Team){
	var temp ListResponse{}
	//将返回的结构体中球队ID找出来
	DB.Table.("response").Find.(&temp)
	var teamA, teamB Team
	//再通过球队ID找出球队的信息
	DB.Table("team").Where("id = ?",temp.TeamAID).Find(&teamA)
	DB.Table("team").Where("id = ?",temp.TeamBID).Find(&teamB)
	return TeamA, TeamB, temp
}