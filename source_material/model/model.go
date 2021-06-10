package model

//User 是用户注册的信息
type User struct {
	ID       int64  `json:"id"`
	Username     string `json:"user_id"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Role     int   `json:"role"`
}

//Player 是球员的信息
type Player struct {
	ID      int
	TeamName    string
	Picture string // 头像
	TeamBelongID  string
}


//球员-球队表
type duiyin struct{
	id int
	TeamAID int 
	playeridd int 
}

//Team 是球队的信息
type Team struct {
	ID int
	Name   string // 队名
	Logo   string // 图标
	//MemberID []*Player
}





// ListResponse ... 返回结构体，获取的球队列表
type ListResponse struct {
	Name        string `json:"name"`
	Date        string `json:"data"`
	Place       string `json: "place"`
	Info        string `json: "info"`
	Appointment  int    `json:"appoinment"`
	TeamAId       int
	TeamBId       int
}





type Appoint struct{
	ID      int
	UserID    int
	GameID int // 头像
	
}
