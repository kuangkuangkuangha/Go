package model

// import "time"

//User 是用户注册的信息
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	ROle  int `json:"role"`
}

//Player 是球员的信息
type Player struct {
	ID      int				`json:"id"`
	Name    string		    `json:"name"`
	Picture string 			`json:"picture"`// 头像	
	TeamBelongID  int	`json:"belongid"`
}


//球员-球队表
type TeamPlayer struct{
	ID int					`json:"id"`
	TeamID int 			`json:"teamid"`   //为什么这个字段在数据库里要以 team_id 的形式存储
	Playerid int 			`json:"playerid"`   //而这个字段要以 playerid 的形式存储， 没有中间的下滑线
//现在数据库中是playerid, 但本应该是player_id ？
}

//Team 是球队的信息
type Team struct {
	ID int					`json:"id"`
	Name   string 			`json:"name"`// 队名	
	Logo   string 			`json:"logo"`// 图标
	MemberID string
}

type Game struct{
	ID int     `json:"id"`
	Name string `jsom:"name"`
	Data string `json:"info"`
	Place string `json:"place"`
	TeamAid int	`json:"teamAid"`
	TeamBid int	`json:"teamBid"`
	Count int  `gorm:"count" json:"count"` 
	// Creat_at time.Time  `gorm:"creat_at"`
}



// ListResponse ... 返回结构体，获取的球赛列表
type ListResponse struct {
	// Name        string `json:"name"`
	Date        string `json:"data"`
	Place       string `json:"place"`//冒号后不能打空格
	// Info        string `json:"info"`

	// Appointment  int    `json:"appoinment"`
	TeamAId       int	`json:"teamAid"`
	TeamBId       int	`json:"teamBid"`
}





type Usergame struct{
	ID      int			`json:"id"`
	UserID    int		`json:"userid"`
	GameID int 			`json:"gameid"`// 头像
	
}
