package handler

//传入球队的id
func Team(id int) []*Team {
	var temp Team
	var ee []int
	//将球队id的对应信息查找出来
	DB.Table("team").where("id = ?", id).Find(&temp)
	//再通过id将该球队的球员查找出来q

	DB.Table("duiyin").where("teamid = ?", id).Find(&ee)

}
